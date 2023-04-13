package dao

import (
	"binghambai.com/shortUrl/app/conn"
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var ctx = context.Background()

func Get(key string) *string {
	val, err := conn.Rd.Get(ctx, key).Result()
	if err != nil {
		panic("connect to redis has error")
	} else if err == redis.Nil {
		log.Println("当前key不存在")
		return nil
	}
	return &val
}

func Exist(key string) bool {
	_, err := conn.Rd.Get(ctx, key).Result()
	if err != nil {
		return false
	}
	return true
}

func Set(key, val string, expire time.Duration) {
	if _, err := conn.Rd.Set(ctx, key, val, expire).Result(); err != nil {
		panic("can not set key-value to redis")
	}
}
