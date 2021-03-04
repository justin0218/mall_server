package services

import (
	"fmt"
	"mall_server/api/proto"
	"mall_server/internal/models/addr"
	"mall_server/internal/models/goods"
	"mall_server/internal/models/order"
	"mall_server/internal/models/sku"
	"mall_server/internal/models/wx"
	"mall_server/pkg/tool"
	"time"
)

//
type OrderService struct {
	baseService
}

//
func (s *OrderService) CreateOrder(in *proto.CreateOrderReq) (ret *proto.CreateOrderRes, err error) {
	db := s.Mysql.Get()
	goodsTbl := goods.NewModel(db)
	goodsInfo, e := goodsTbl.GetOneById(int(in.GoodsId))
	if e != nil || goodsInfo.Id == 0 {
		err = fmt.Errorf("商品不存在")
		return
	}
	if goodsInfo.Status != 1 {
		err = fmt.Errorf("商品暂未上架")
		return
	}
	orderCode := fmt.Sprintf("%d", tool.GetSnowflakeId())
	tx := db.Begin()
	skuTbl := sku.NewModel(tx)
	err = skuTbl.DecrInventory(int(in.SkuId), int(in.BuyNum))
	if err != nil {
		tx.Rollback()
		return
	}
	skuInfo, e := skuTbl.GetById(int(in.SkuId))
	if e != nil || skuInfo.Id == 0 {
		err = fmt.Errorf("sku不存在")
		return
	}
	orderTbl := order.NewModel(tx)
	inOrder := order.Order{
		OrderCode: orderCode,
		GoodsId:   in.GoodsId,
		GoodsName: goodsInfo.Name,
		SkuId:     in.SkuId,
		SkuName:   skuInfo.Name,
		Count:     in.BuyNum,
		Price:     int64(skuInfo.Price) * in.BuyNum,
		Uid:       in.Uid,
	}
	err = orderTbl.Create(inOrder)
	if err != nil {
		tx.Rollback()
		return
	}
	addrTbl := addr.NewModel(tx)
	err = addrTbl.Create(addr.Addr{
		Uid:       in.Uid,
		OrderCode: orderCode,
		Phone:     in.Phone,
		Name:      in.Name,
		Province:  in.Province,
		City:      in.City,
		Region:    in.Region,
		Addr:      in.Addr,
	})
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	if err != nil {
		return
	}
	ret = new(proto.CreateOrderRes)
	ret.OrderCode = orderCode
	ret.Price = inOrder.Price
	ret.GoodsName = goodsInfo.Name
	ret.SkuName = skuInfo.Name
	return
}

func (s *OrderService) AutoCloseOrder(in *proto.AutoCloseOrderReq) (ret *proto.AutoCloseOrderRes, err error) {
	db := s.Mysql.Get()
	orderTbl := order.NewModel(db)
	orders, e := orderTbl.GetByStatus(0)
	if e != nil {
		e = err
		return
	}
	for _, v := range orders {
		if time.Now().Unix()-v.CreatedAt.Unix() >= 600 {
			err = orderTbl.UpdateStatusByOrderCode(v.OrderCode, -1)
			if err != nil {
				continue
			}
		}
	}
	ret = new(proto.AutoCloseOrderRes)
	return
}

func (s *OrderService) ConfirmOrder(in *wx.WxpayReq) error {
	db := s.Mysql.Get()
	orderTbl := order.NewModel(db)
	return orderTbl.UpdateByOrderCode(in)
}
