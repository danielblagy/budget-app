package access

import (
	"context"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/danielblagy/budget-app/internal/service/cache"
	"github.com/danielblagy/budget-app/internal/service/users"
)

type Service interface {
	LogIn(ctx context.Context, login *model.Login) (*model.UserTokens, error)
	// LogOut adds jwt tokens to blacklist.
	LogOut(ctx context.Context, accessToken, refreshToken string) error
	// Authorize returns username if successfully authenticated.
	Authorize(ctx context.Context, token string) (string, error)
}

type service struct {
	usersService users.Service
	cacheService cache.Service
}

func NewService(usersService users.Service, cacheService cache.Service) Service {
	return &service{
		usersService: usersService,
		cacheService: cacheService,
	}
}
