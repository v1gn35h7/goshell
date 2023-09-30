package redis

import (
	"github.com/redis/go-redis"
	"github.com/redis/go-redis/v9"
	"github.com/v1gn35h7/goshell/pkg/redis"
)

func NewClient() redis.Client {

	// Setup redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return *rdb
}
