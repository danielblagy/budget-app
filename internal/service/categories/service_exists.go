package categories

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (s service) Exists(ctx context.Context, username string, categoryID int64) (bool, error) {
	_, err := s.categoriesQuery.Get(ctx, username, categoryID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, errors.Wrap(err, "can't check if category exists")
	}

	return true, nil
}
