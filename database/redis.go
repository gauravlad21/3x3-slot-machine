package database

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/slot_machine/config"
)

var redisClient *redis.Client

func GetRedisClient() *redis.Client {
	if redisClient == nil {
		InitRedis(config.AppConfig.RedisURI)
	}
	return redisClient
}

func InitRedis(redisURI string) {
	redisClient = redis.NewClient(&redis.Options{
		Addr: redisURI,
	})

	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis")
}
