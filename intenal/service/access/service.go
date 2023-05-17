package access

import (
	"context"

	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/danielblagy/budget-app/intenal/service/users"
)

type Service interface {
	LogIn(ctx context.Context, login *model.Login) (*model.UserTokens, error)
	// Authorize returns username if successfully authenticated.
	Authorize(ctx context.Context, token string) (string, error)
}

type service struct {
	usersService users.Service
}

func NewService(usersService users.Service) Service {
	return &service{
		usersService: usersService,
	}
}
