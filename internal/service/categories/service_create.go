package categories

import (
	"context"
	"fmt"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/pkg/errors"
)

func (s service) Create(ctx context.Context, username string, category *model.CreateCategory) (*model.Category, error) {
	createdCategory, err := s.categoriesQuery.Add(ctx, username, category)
	if err != nil {
		return nil, errors.Wrap(err, "can't create category")
	}

	err = s.cacheService.Set(ctx, fmt.Sprintf("%s:category:%d", username, createdCategory.ID), createdCategory)
	if err != nil {
		s.logger.Error("can't add to created category to cache", "err", err.Error())
	}

	return createdCategory, nil
}
