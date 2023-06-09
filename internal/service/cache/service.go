package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Service interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) ([]byte, bool, error)
	SetWithExpiration(ctx context.Context, key string, value interface{}, expiration time.Duration) error
}

type service struct {
	redisClient *redis.Client
}

func NewService(redisClient *redis.Client) Service {
	return &service{
		redisClient: redisClient,
	}
}
