package store

import (
	"context"
	"fmt"
	"os"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	storeService = &StorageService{}
	ctx = context.Background()
)

const CacheDuration = 6 * time.Hour


type StorageService struct {
	redisClient *redis.Client
}


func InitializeStore() *StorageService {

	redisURL := os.Getenv("REDIS_URL")

	var redisClient *redis.Client

	if redisURL != "" {

		opt, err := redis.ParseURL(redisURL)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse REDIS_URL: %v", err))
		}

		redisClient = redis.NewClient(opt)

	} else {

		// fallback local
		redisClient = redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			Password: "",
			DB: 0,
		})

	}

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong = %s\n", pong)

	storeService.redisClient = redisClient
	return storeService
}

func SaveUrlMapping(shortUrl string, originalUrl string, userId string){
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}

}

func RetrieveInitialUrl(shortUrl string) string{
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}