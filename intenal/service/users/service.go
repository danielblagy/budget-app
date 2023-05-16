package users

import (
	"context"

	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/jackc/pgx/v5"
)

type Service interface {
	// GetUsers retuns users with omitted Password field (empty string)
	GetUsers(ctx context.Context) ([]*model.User, error)
	// GetUser retuns user with omitted Password field (empty string)
	GetUser(ctx context.Context, username string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	Exists(ctx context.Context, username string) (bool, error)
	UserWithEmailExists(ctx context.Context, email string) (bool, error)
}

type service struct {
	db *pgx.Conn
}

func NewService(db *pgx.Conn) Service {
	return &service{
		db: db,
	}
}
