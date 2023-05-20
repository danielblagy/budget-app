package categories

import (
	"context"

	"github.com/danielblagy/budget-app/intenal/model"
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

	return updatedCategory, nil
}
