package operation

import (
	"github.com/go-redis/redis"
	"fmt"
)

func HashOperation(client *redis.Client) {
	client.HSet("user_ray", "name", "ray")
	client.HSet("user_ray", "age", "18")

	client.HMSet("user_dong", map[string]interface{}{
		"name": "dong",
		"age":  "19",})

	fields, err := client.HMGet("user_dong", "name", "age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("fields in user_dong: ", fields)

	length, err := client.HLen("user_ray").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("field count in user_ray: ", length)

	client.HDel("user_ray", "age")
	age, err := client.HGet("user_ray", "age").Result()
	if err != nil {
		fmt.Printf("Get user_ray age err: %v\n", err.Error())
	} else {
		fmt.Println("user_ray age is: ", age)
	}
}
