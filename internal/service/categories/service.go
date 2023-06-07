package categories

import (
	"context"

	"github.com/danielblagy/budget-app/internal/db"
	"github.com/danielblagy/budget-app/internal/model"
)

type Service interface {
	// GetAll returns all users's categories.
	GetAll(ctx context.Context, username string, categoryType model.CategoryType) ([]*model.Category, error)
	Get(ctx context.Context, username string, categoryID int64) (*model.Category, error)
	Create(ctx context.Context, username string, category *model.CreateCategory) (*model.Category, error)
	Update(ctx context.Context, username string, updateData *model.UpdateCategory) (*model.Category, error)
	Delete(ctx context.Context, username string, categoryID int64) (*model.Category, error)
}

type service struct {
	categoriesQuery db.CategoriesQuery
}

func NewService(categoriesQuery db.CategoriesQuery) Service {
	return &service{
		categoriesQuery: categoriesQuery,
	}
}
