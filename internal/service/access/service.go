package access

import (
	"context"

	"github.com/danielblagy/budget-app/internal/model"
	persistent_store "github.com/danielblagy/budget-app/internal/service/persistent-store"
	"github.com/danielblagy/budget-app/internal/service/users"
)

type Service interface {
	LogIn(ctx context.Context, login *model.Login) (*model.UserTokens, error)
	// LogOut adds jwt tokens to blacklist.
	LogOut(ctx context.Context, accessToken, refreshToken string) error
	// Authorize returns username if successfully authenticated.
	Authorize(ctx context.Context, token string) (string, error)
	// Refresh adds jwt tokens to blacklist and returns newely generated token pair.
	Refresh(ctx context.Context, accessToken, refreshToken string) (*model.UserTokens, error)
}

type service struct {
	usersService           users.Service
	persistentStoreService persistent_store.Service
}

func NewService(usersService users.Service, persistentStoreService persistent_store.Service) Service {
	return &service{
		usersService:           usersService,
		persistentStoreService: persistentStoreService,
	}
}
