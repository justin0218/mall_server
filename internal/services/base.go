package services

import (
	"mall_server/store"
)

type baseService struct {
	Mysql  store.Mysql
	Redis  store.Redis
	Config store.Config
}
