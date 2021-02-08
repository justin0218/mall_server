package api

import (
	"context"
	"mall_server/api/proto"
)

func (s *MallSvr) GetGoodsDetail(ctx context.Context, req *proto.GetGoodsDetailReq) (ret *proto.GetGoodsDetailRes, err error) {
	return s.GoodsService.GetDetailById(int(req.GoodsId))
}
