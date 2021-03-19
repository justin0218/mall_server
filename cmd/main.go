package main

import (
	"fmt"
	"mall_server/api"
	"mall_server/pkg/rabbitmq/auto_close_order"
	"mall_server/pkg/rabbitmq/consume_wx_pay_notice"
	"mall_server/store"
	"time"
)

func main() {
	mq := new(store.Rabbitmq)
	mq.Get()
	redis := new(store.Redis)
	err := redis.Get().Ping().Err()
	if err != nil {
		panic(err)
	}
	mysql := new(store.Mysql)
	mysql.Get()
	log := new(store.Log)
	//go etcd.Register()
	go api.GrpcServer()
	log.Get().Debug("server started at %v", time.Now())
	fmt.Printf("server started at %v \n", time.Now())
	go consume_wx_pay_notice.Do()
	go auto_close_order.Do()
	select {}
}
