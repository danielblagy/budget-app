package categories

import (
	"context"

	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("category not found")

func (s service) Get(ctx context.Context, username string, categoryID int64) (*model.Category, error) {
	category, err := s.categoriesQuery.Get(ctx, username, categoryID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, errors.Wrap(err, "can't get category")
	}

	return category, nil
}
