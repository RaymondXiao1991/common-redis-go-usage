package operation

import (
	"testing"
	"redis-go"
)

func TestHashOperation(t *testing.T) {
	client := redis_go.NewRedisClient()
	HashOperation(client)
}
