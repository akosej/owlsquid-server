package system

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var CTX context.Context
var RDB *redis.Client

func ConnectMysql() {
	connection, err := gorm.Open(mysql.Open(MysqlCredentials), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	DB = connection
	connection.AutoMigrate(&NavigationUsers{})
}

func ConnectRedis() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     RedisServer,
		Password: RedisServerPass,
	})
	CTX = ctx
	RDB = rdb
}
