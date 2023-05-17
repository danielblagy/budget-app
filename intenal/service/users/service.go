package users

import (
	"context"

	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/jackc/pgx/v5"
)

type Service interface {
	// GetAll retuns users with omitted Password field (empty string)
	GetAll(ctx context.Context) ([]*model.User, error)
	// Get retuns user with omitted Password field (empty string)
	Get(ctx context.Context, username string) (*model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
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
