package categories

import (
	"context"
	"fmt"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

func (s service) Delete(ctx context.Context, username string, categoryID int64, deleteEntries bool) (*model.Category, error) {
	tx, err := s.queryFactory.GetDbRunner().Begin(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "can't begin tx")
	}

	deletedCategory, err := s.delete(ctx, tx, username, categoryID, deleteEntries)
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			err = multierr.Append(err, rollbackErr)
		}
		return nil, errors.Wrap(err, "can't delete categories")
	}

	if commitErr := tx.Commit(ctx); commitErr != nil {
		return nil, errors.Wrap(commitErr, "can't commit tx")
	}

	return deletedCategory, nil
}

func (s service) delete(ctx context.Context, tx pgx.Tx, username string, categoryID int64, deleteEntries bool) (*model.Category, error) {
	categoriesQuery := s.queryFactory.NewCategoriesQuery(tx)

	deletedCategory, err := categoriesQuery.Delete(ctx, username, categoryID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, errors.Wrap(err, "can't delete category")
	}

	err = s.cacheService.Delete(ctx, fmt.Sprintf("%s:category:%d", username, categoryID))
	if err != nil {
		return nil, errors.Wrap(err, "can't delete category from cache")
	}

	entriesQuery := s.queryFactory.NewEntriesQuery(tx)

	if deleteEntries {
		err = entriesQuery.DeleteByUserAndCategory(ctx, username, categoryID)
		if err != nil {
			return nil, errors.Wrap(err, "can't delete category entries")
		}

		return deletedCategory, nil
	}

	err = entriesQuery.SetNullCategory(ctx, username, categoryID)
	if err != nil {
		return nil, errors.Wrap(err, "can't set category entries's category_id to null")
	}

	return deletedCategory, nil
}
