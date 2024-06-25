package configs

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client
var ctx = context.Background()

func ConnectRedis() {
    RedisClient = redis.NewClient(&redis.Options{
        Addr:     os.Getenv("REDIS_HOST"),
        Password: os.Getenv("REDIS_PASSWORD"), 
        DB:       0,                          
    })

    _, err := RedisClient.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("Could not connect to Redis: %v", err)
    }
    log.Println("Connected to Redis")
}

func GetRedisClient() *redis.Client {
    return RedisClient
}

func GetRedisContext() context.Context {
    return ctx
}
