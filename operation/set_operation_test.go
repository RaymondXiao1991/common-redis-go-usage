package operation

import (
	"testing"
	"redis-go"
)

func TestSetOperation(t *testing.T) {
	client := redis_go.NewRedisClient()
	SetOperation(client)
}
