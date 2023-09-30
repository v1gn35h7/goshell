package redis

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/redis/go-redis/v9"
	"github.com/v1gn35h7/goshell/pkg/redis"
)

func Get(client redis.Client, key string, logger logr.Logger) (string, error) {
	ctx := context.Background()

	res, err := client.Get(ctx, key).Result()

	if err != nil {
		logger.Error(err, "Falied to read key")
		return "", err
	}

	return res, nil
}
