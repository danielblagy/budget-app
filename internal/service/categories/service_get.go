package categories

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("category not found")

func (s service) Get(ctx context.Context, username string, categoryID int64) (*model.Category, error) {
	cacheKey := fmt.Sprintf("%s:category:%d", username, categoryID)
	cacheValueBytes, ok, err := s.cacheService.Get(ctx, cacheKey)
	if err != nil {
		return nil, errors.Wrap(err, "can't get category from cache")
	}
	if ok {
		var category *model.Category
		if unmarshalErr := json.Unmarshal(cacheValueBytes, &category); unmarshalErr != nil {
			return nil, errors.Wrap(unmarshalErr, "can't unmarshal data")
		}

		return category, nil
	}

	category, err := s.categoriesQuery.Get(ctx, username, categoryID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, errors.Wrap(err, "can't get category")
	}

	err = s.cacheService.Set(ctx, cacheKey, category)
	if err != nil {
		return nil, err
	}

	return category, nil
}
