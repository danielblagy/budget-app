package cache

import (
	"context"

	"github.com/pkg/errors"
)

func (s service) Delete(ctx context.Context, key string) error {
	if err := s.redisClient.Del(ctx, key).Err(); err != nil {
		return errors.Wrap(err, "can't delete cache value")
	}

	return nil
}
