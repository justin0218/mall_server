package auth_server

import (
	"mall_server/pkg/etcd"
	"sync"
)

var once sync.Once
var conn AuthClient

func GetClient() AuthClient {
	once.Do(func() {
		conn = NewAuthClient(etcd.Discovery("auth_server"))
	})
	return conn
}
