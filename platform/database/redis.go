package database

import (
	"github.com/go-redis/redis/v7"
	"os"
)

func RedisConnection() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Username: os.Getenv("REDIS_SERVER_NAME"),
		Password: os.Getenv("REDIS_SERVER_PASS"),
		Addr:     os.Getenv("REDIS_SERVER_URL"),
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}
