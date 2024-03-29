package categories

import (
	"context"

	"github.com/danielblagy/budget-app/internal/db"
	"github.com/danielblagy/budget-app/internal/model"
	"github.com/danielblagy/budget-app/internal/service/cache"
	log "github.com/inconshreveable/log15"
)

type Service interface {
	// GetAll returns all users's categories.
	GetAll(ctx context.Context, username string, categoryType model.CategoryType) ([]*model.Category, error)
	Get(ctx context.Context, username string, categoryID int64) (*model.Category, error)
	Create(ctx context.Context, username string, category *model.CreateCategory) (*model.Category, error)
	Update(ctx context.Context, username string, updateData *model.UpdateCategory) (*model.Category, error)
	Delete(ctx context.Context, username string, categoryID int64, deleteEntries bool) (*model.Category, error)
	Exists(ctx context.Context, username string, categoryID int64) (bool, error)
}

type service struct {
	logger          log.Logger
	categoriesQuery db.CategoriesQuery
	cacheService    cache.Service
	queryFactory    db.QueryFactory
}

func NewService(
	logger log.Logger,
	categoriesQuery db.CategoriesQuery,
	cacheService cache.Service,
	queryFactory db.QueryFactory,
) Service {
	return &service{
		logger:          logger,
		categoriesQuery: categoriesQuery,
		cacheService:    cacheService,
		queryFactory:    queryFactory,
	}
}
