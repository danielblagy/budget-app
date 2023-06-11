package persistent_store

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

func (s service) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if err := s.redisClient.Set(ctx, key, value, expiration).Err(); err != nil {
		return errors.Wrap(err, "can't set persistent store value")
	}

	return nil
}
