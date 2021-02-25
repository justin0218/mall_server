package services

import (
	"mall_server/api/proto"
	"mall_server/internal/models/goods"
	"mall_server/internal/models/sku"
)

//
type GoodsService struct {
	baseService
}

//
func (s *GoodsService) GetDetailById(goodsId int) (ret *proto.GetGoodsDetailRes, err error) {
	db := s.Mysql.Get()
	goodsInfo, e := goods.NewModel(db).GetOneById(goodsId)
	if e != nil {
		err = e
		return
	}
	ret = new(proto.GetGoodsDetailRes)
	ret.Name = goodsInfo.Name
	ret.Images = goodsInfo.Images
	ret.GoodsId = int64(goodsInfo.Id)
	ret.Status = int32(goodsInfo.Status)
	//查询商品下sku
	skuInfo, e := sku.NewModel(db).GetMultiByGoodsId(goodsInfo.Id)
	if e != nil {
		err = e
		return
	}
	for _, v := range skuInfo {
		if v.Inventory == 0 {
			continue
		}
		ret.Skus = append(ret.Skus, &proto.Sku{
			Id:        int64(v.Id),
			Name:      v.Name,
			Cover:     v.Cover,
			GoodsId:   int64(v.GoodsId),
			Price:     int64(v.Price),
			Inventory: int64(v.Inventory),
		})
	}
	return
}
