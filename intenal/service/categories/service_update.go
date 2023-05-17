package categories

import (
	"context"
	"fmt"

	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (s service) Update(ctx context.Context, id int64, category *model.NewCategory) (*model.Category, error) {
	var updatedCategory model.Category
	err := pgxscan.Get(ctx, s.db, &updatedCategory, fmt.Sprintf("update categories set name = '%s' where id = '%d' returning id, user_id, name", category.Name, id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, errors.Wrap(err, "can't update category")
	}

	return &updatedCategory, nil
}
