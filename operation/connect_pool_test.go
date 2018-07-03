package operation

import (
	"testing"
	"redis-go"
)

func TestConnectPool(t *testing.T) {
	client := redis_go.NewRedisClient()
	ConnectPool(client)
}
