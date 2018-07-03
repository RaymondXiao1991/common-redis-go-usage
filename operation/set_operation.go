package operation

import (
	"github.com/go-redis/redis"
	"fmt"
)

func SetOperation(client *redis.Client) {
	client.SAdd("blacklist", "Obama") // 向set中添加元素
	client.SAdd("blacklist", "Hillary")
	client.SAdd("blacklist", "Chump")

	client.SAdd("whitelist", "Chump")

	// 判断元素是否在集合中
	isMember, err := client.SIsMember("blacklist", "Bush").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Is Bush in blacklist?: ", isMember)

	// 求交集
	names, err := client.SInter("blacklist", "whitelist").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Inter result: ", names)

	// 获取指定集合中的所有元素
	names, err = client.SMembers("blacklist").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("All members of blacklist: ", names)

}
