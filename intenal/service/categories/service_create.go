package categories

import (
	"context"

	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/pkg/errors"
)

func (s service) Create(ctx context.Context, username string, category *model.CreateCategory) (*model.Category, error) {
	createdCategory, err := s.categoriesQuery.Add(ctx, username, category.Name)
	if err != nil {
		return nil, errors.Wrap(err, "can't create category")
	}

	return createdCategory, nil
}
