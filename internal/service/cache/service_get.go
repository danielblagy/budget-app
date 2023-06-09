package cache

import (
	"context"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

func (s service) Get(ctx context.Context, key string) (interface{}, bool, error) {
	result := s.redisClient.Get(ctx, key)
	if err := result.Err(); err != nil {
		if err == redis.Nil {
			return nil, false, nil
		}

		return nil, false, errors.Wrap(err, "can't get cache value")
	}

	var value interface{}
	if err := result.Scan(value); err != nil {
		return nil, false, errors.Wrap(err, "can't scan cache value")
	}

	return value, true, nil
}
