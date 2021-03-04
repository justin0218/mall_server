package store

import (
	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/streadway/amqp"
	"sync"
)

var (
	logOnce      sync.Once
	redisOnce    sync.Once
	configOnce   sync.Once
	mysqlOnce    sync.Once
	rabbitMqOnce sync.Once
)

var (
	logClient      *logs.BeeLogger
	redisClient    *redis.Client
	config         cfg
	mysqlClient    *gorm.DB
	rabbitMqClient *amqp.Connection
)
