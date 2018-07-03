package operation

import (
	"github.com/go-redis/redis"
	"sync"
	"fmt"
)

func ConnectPool(client *redis.Client) {
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < 100; j++ {
				client.Set(fmt.Sprintf("name%d", j), fmt.Sprintf("ray%d", j), 0).Err()
				client.Get(fmt.Sprintf("name%d", j)).Result()
			}

			fmt.Printf("PoolStatus, TotalConns: %d, FreeConns: %d\n", client.PoolStats().TotalConns, client.PoolStats().FreeConns)
		}()
	}

	wg.Wait()
}
