package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var (
	Rdb *redis.Client
	Ctx = context.Background()
)

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // adjust if Redis is on another host or port
		Password: "",               // no password by default
		DB:       0,                // default DB
	})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}
