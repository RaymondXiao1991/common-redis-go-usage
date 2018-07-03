package redis_go

import (
	"github.com/go-redis/redis"
	"fmt"
)

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 5,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}
