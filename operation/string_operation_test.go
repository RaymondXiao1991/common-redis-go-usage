package operation

import (
	"testing"
	"redis-go"
)

func TestStringOperation(t *testing.T) {
	client := redis_go.NewRedisClient()
	StringOperation(client)
}
