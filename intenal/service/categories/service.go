package categories

import (
	"context"

	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/jackc/pgx/v5"
)

type Service interface {
	// GetAll returns all users's categories.
	GetAll(ctx context.Context, username string) ([]*model.Category, error)
	Get(ctx context.Context, username string, categoryID int64) (*model.Category, error)
	Create(ctx context.Context, username string, category *model.CreateCategory) (*model.Category, error)
	Update(ctx context.Context, username string, updateData *model.UpdateCategory) (*model.Category, error)
}

type service struct {
	db *pgx.Conn
}

func NewService(db *pgx.Conn) Service {
	return &service{
		db: db,
	}
}
