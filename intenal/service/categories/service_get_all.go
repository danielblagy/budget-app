package categories

import (
	"context"
	"fmt"

	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/pkg/errors"
)

func (s service) GetAll(ctx context.Context, username string) ([]*model.Category, error) {
	var categories []*model.Category
	err := pgxscan.Select(ctx, s.db, &categories, fmt.Sprintf("select id, user_id, name from categories where user_id = '%s'", username))
	if err != nil {
		return nil, errors.Wrap(err, "can't get categories from db")
	}

	return categories, nil
}
