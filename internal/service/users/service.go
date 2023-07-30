package users

import (
	"context"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=Service --case=underscore

type Service interface {
	// GetAll retuns users with omitted Password field (empty string)
	GetAll(ctx context.Context) ([]*model.User, error)
	// Get retuns user with omitted Password field (empty string)
	Get(ctx context.Context, username string) (*model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Exists(ctx context.Context, username string) (bool, error)
	UserWithEmailExists(ctx context.Context, email string) (bool, error)
	GetPasswordHash(ctx context.Context, username string) (string, error)
}

type service struct {
	db *pgxpool.Pool
}

func NewService(db *pgxpool.Pool) Service {
	return &service{
		db: db,
	}
}
