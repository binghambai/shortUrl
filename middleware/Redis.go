package middleware

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var Rd *redis.Client
var ctx = context.Background()

type RedisCustom struct {
	redisClient *redis.Client
	ctx         context.Context
}

func InitRedis() {
	Rd = redis.NewClient(&redis.Options{
		Addr:     "localhost:56379",
		Password: "redis",
	})
}

func RedisGet(key string) (string, error) {
	return Rd.Get(ctx, key).Result()
}
