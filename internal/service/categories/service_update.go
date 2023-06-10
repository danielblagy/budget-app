package categories

import (
	"context"
	"fmt"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (s service) Update(ctx context.Context, username string, updateData *model.UpdateCategory) (*model.Category, error) {
	updatedCategory, err := s.categoriesQuery.Update(ctx, username, updateData.ID, updateData.Name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, errors.Wrap(err, "can't update category")
	}

	err = s.cacheService.Set(ctx, fmt.Sprintf("%s:category:%d", username, updatedCategory.ID), updatedCategory)
	if err != nil {
		return nil, errors.Wrap(err, "can't update category cache")
	}

	return updatedCategory, nil
}
