package operation

import (
	"github.com/go-redis/redis"
	"fmt"
)

func ListOperation(client *redis.Client) {
	client.RPush("fruit", "apple")
	client.LPush("fruit", "banana")

	length, err := client.LLen("fruit").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("length: ", length)

	value, err := client.LPop("fruit").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("fruit: ", value)

	value, err = client.RPop("fruit").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("fruit: ", value)
}
