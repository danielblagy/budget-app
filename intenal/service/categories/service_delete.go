package categories

import (
	"context"
	"errors"
	"fmt"

	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

func (s service) Delete(ctx context.Context, username string, categoryID int64) (*model.Category, error) {
	var deletedCategory model.Category
	err := pgxscan.Get(ctx, s.db, &deletedCategory, fmt.Sprintf("delete from categories where user_id = '%s' and id = '%d' returning id, user_id, name", username, categoryID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &deletedCategory, nil
}
