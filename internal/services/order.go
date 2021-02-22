package services

import (
	"fmt"
	"mall_server/api/proto"
	"mall_server/internal/models/addr"
	"mall_server/internal/models/goods"
	"mall_server/internal/models/order"
	"mall_server/internal/models/sku"
	"mall_server/pkg/tool"
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
	err = orderTbl.Create(order.Order{
		OrderCode: orderCode,
		GoodsId:   in.GoodsId,
		GoodsName: goodsInfo.Name,
		SkuId:     in.SkuId,
		SkuName:   skuInfo.Name,
		Count:     in.BuyNum,
		Price:     skuInfo.Price * int(in.BuyNum),
		Uid:       in.Uid,
	})
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
	return
}