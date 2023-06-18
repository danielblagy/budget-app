package persistent_store

import (
	"context"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

func (s service) Get(ctx context.Context, key string) ([]byte, bool, error) {
	result := s.redisClient.Get(ctx, key)
	if err := result.Err(); err != nil {
		if err == redis.Nil {
			return nil, false, nil
		}

		return nil, false, errors.Wrap(err, "can't get persistent store value")
	}

	valueBytes, err := result.Bytes()
	if err != nil {
		return nil, false, errors.Wrap(err, "can't convert cache value to []byte")
	}

	return valueBytes, true, nil
}
