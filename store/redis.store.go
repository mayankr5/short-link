package store

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/mayankr5/url_shortner/config"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
)

func InitializeStore() (*StorageService, error) {
	addr := config.Config("REDISHOST") + ":" + config.Config("REDISPORT")

	if addr == ":" {
		addr = "localhost:6379"
	}

	pass := config.Config("REDISPASSWORD")

	if pass == "" {
		pass = ""
	}
	fmt.Println(addr, " ", pass)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		fmt.Printf("Error init Redis: %v", err)
		return nil, err
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s} \n", pong)
	storeService.redisClient = redisClient
	return storeService, err
}

func SaveUrlMapping(shortUrl string, originalUrl string, userId uuid.UUID, expirationDate time.Time) error {
	CacheDuration := time.Until(expirationDate)
	err := storeService.redisClient.Set(shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		fmt.Printf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl)
		return err
	}
	return nil
}

func RetrieveInitialUrl(shortUrl string) (string, error) {
	result, err := storeService.redisClient.Get(shortUrl).Result()
	if err != nil {
		fmt.Printf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl)
		return "", err
	}
	return result, nil
}
