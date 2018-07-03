package operation

import (
	"testing"
	"redis-go"
)

func TestListOperation(t *testing.T) {
	client := redis_go.NewRedisClient()
	ListOperation(client)
}
