package access

import (
	"context"

	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/danielblagy/budget-app/intenal/service/users"
	"github.com/jackc/pgx/v5"
)

type Service interface {
	LogIn(ctx context.Context, login *model.Login) (*model.UserTokens, error)
	// Authenticate returns true if successfully authenticated.
	Authenticate(ctx context.Context, token string) (bool, error)
}

type service struct {
	db           *pgx.Conn
	usersService users.Service
}

func NewService(db *pgx.Conn, usersService users.Service) Service {
	return &service{
		db:           db,
		usersService: usersService,
	}
}
