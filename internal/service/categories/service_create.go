package categories

import (
	"context"
	"fmt"
	"log"

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
		// TODO log the error with custom logger
		log.Println("[Error]", err)
	}

	return createdCategory, nil
}
