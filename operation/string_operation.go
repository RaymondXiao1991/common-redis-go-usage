package operation

import (
	"github.com/go-redis/redis"
	"fmt"
	"time"
)

// StringOperation  string操作
func StringOperation(client *redis.Client) {
	// expiration 设置过期时间
	err := client.Set("sys", "e-bank", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("sys").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("sys", val)

	err = client.Set("age", 20, 1*time.Second).Err()
	if err != nil {
		panic(err)
	}

	client.Incr("age") // 自增
	client.Incr("age") // 自增
	client.Decr("age") // 自减

	val, err = client.Get("age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("age", val)

	time.Sleep(2 * time.Second)
	val2, err := client.Get("age").Result()
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
	}
	fmt.Println("age", val2)
}
