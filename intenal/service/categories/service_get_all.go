package categories

import (
	"context"

	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/pkg/errors"
)

func (s service) GetAll(ctx context.Context, username string) ([]*model.Category, error) {
	categories, err := s.categoriesQuery.GetAll(ctx, username)
	if err != nil {
		return nil, errors.Wrap(err, "can't get categories")
	}

	return categories, nil
}
