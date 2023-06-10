package categories

import (
	"context"
	"fmt"
	"log"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (s service) Delete(ctx context.Context, username string, categoryID int64) (*model.Category, error) {
	deletedCategory, err := s.categoriesQuery.Delete(ctx, username, categoryID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, errors.Wrap(err, "can't delete category")
	}

	err = s.cacheService.Delete(ctx, fmt.Sprintf("%s:category:%d", username, categoryID))
	if err != nil {
		// TODO log the error with custom logger
		log.Println("[Error]", err)
	}

	return deletedCategory, nil
}
