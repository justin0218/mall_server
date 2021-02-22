package api

import (
	"context"
	"mall_server/api/proto"
)

func (s *MallSvr) CreateOrder(ctx context.Context, req *proto.CreateOrderReq) (ret *proto.CreateOrderRes, err error) {
	return s.OrderService.CreateOrder(req)
}
