package users

import (
	"context"

	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/jackc/pgx/v5"
)

type Service interface {
	GetUsers(ctx context.Context) ([]*model.User, error)
	GetUser(ctx context.Context, username string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	Exists(ctx context.Context, username string) (bool, error)
}

type service struct {
	db *pgx.Conn
}

func NewService(db *pgx.Conn) Service {
	return &service{
		db: db,
	}
}
