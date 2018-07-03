package main

import (
	"fmt"
	"github.com/go-redis/redis"
	rgo "redis-go"
)

func main() {
	client := rgo.NewRedisClient()

	err := client.Set("cpn", "bayes", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("cpn").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("cpn", val)

	val2, err := client.Get("cpn2").Result()
	if err == redis.Nil {
		fmt.Println("cpn2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("cpn2", val2)
	}

	return
}
