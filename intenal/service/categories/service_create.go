package categories

import (
	"context"
	"fmt"

	"github.com/danielblagy/budget-app/intenal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/pkg/errors"
)

func (s service) Create(ctx context.Context, username string, category *model.NewCategory) (*model.Category, error) {
	rows, err := s.db.Query(ctx, fmt.Sprintf("insert into categories (user_id, name) values ('%s', '%s') returning id, user_id, name", username, category.Name))
	if err != nil {
		return nil, errors.Wrap(err, "can't insert into db")
	}

	var createdCategory model.Category
	err = pgxscan.ScanOne(&createdCategory, rows)
	if err != nil {
		return nil, errors.Wrap(err, "can't return inserted entity")
	}

	return &createdCategory, nil
}
