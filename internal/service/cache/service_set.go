package cache

import (
	"context"

	"github.com/pkg/errors"
)

func (s service) Set(ctx context.Context, key string, value interface{}) error {
	if err := s.redisClient.Set(ctx, key, value, 0).Err(); err != nil {
		return errors.Wrap(err, "can't set cache value")
	}

	return nil
}
