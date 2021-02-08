package api

import (
	"fmt"
	"google.golang.org/grpc"
	"mall_server/api/proto"
	"mall_server/internal/services"
	"mall_server/store"
	"net"
)

type MallSvr struct {
	GoodsService services.GoodsService
}

func GrpcServer() {
	conf := new(store.Config)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s", conf.Get().Etcd.Key))
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	svr := grpc.NewServer(opts...)
	proto.RegisterMallServer(svr, &MallSvr{})
	err = svr.Serve(lis)
	if err != nil {
		panic(err)
	}
}
