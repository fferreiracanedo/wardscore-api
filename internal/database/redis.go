package database

import (
	"context"
	"log"
	"time"
	"wardscore-api/internal/config"

	"github.com/redis/go-redis/v9"
)


var RedisClient *redis.Client
var ctx = context.Background()


func ConnectRedis() {
	opt, err := redis.ParseURL(config.AppConfig.RedisURL)
	if err != nil {
		log.Fatal("❌ Falha ao parsear URL do Redis:", err)
	}

	RedisClient = redis.NewClient(opt)


	_, err = RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("❌ Falha ao conectar com Redis:", err)
	}

	log.Println("✅ Conectado ao Redis")
}


func SetCache(key string, value interface{}, expiration time.Duration) error {
	return RedisClient.Set(ctx, key, value, expiration).Err()
}

func GetCache(key string) (string, error) {
    return RedisClient.Get(ctx, key).Result()
}

func DeleteCache(key string) error {
    return RedisClient.Del(ctx, key).Err()
}

func ExistsCache(key string) (bool, error) {
    result, err := RedisClient.Exists(ctx, key).Result()
    return result > 0, err
}