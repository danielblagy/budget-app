package categories

import (
	"context"
	"fmt"

	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("category not found")

func (s service) Get(ctx context.Context, categoryID int64) (*model.Category, error) {
	var category model.Category
	err := pgxscan.Get(ctx, s.db, &category, fmt.Sprintf("select id, user_id, name from categories where id = '%d'", categoryID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, errors.Wrap(err, "can't get category from db")
	}

	return &category, nil
}
