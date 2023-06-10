package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=Service --case=underscore

type Service interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) ([]byte, bool, error)
	SetWithExpiration(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Delete(ctx context.Context, key string) error
}

type service struct {
	redisClient *redis.Client
}

func NewService(redisClient *redis.Client) Service {
	return &service{
		redisClient: redisClient,
	}
}
