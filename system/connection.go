package system

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var CTX context.Context
var RDB *redis.Client

// ConnectServerRedis -- Connect to the redis server
func ConnectServerRedis() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     RedisServer,
		Password: RedisServerPass,
	})
	CTX = ctx
	RDB = rdb
}
