package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var RDB *redis.Client

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
}

func GetCached(key string) (string, error) {
	return RDB.Get(ctx, key).Result()
}

func SetCached(key, val string, ttl time.Duration) {
	RDB.Set(ctx, key, val, ttl)
}
