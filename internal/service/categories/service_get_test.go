package categories

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	dbMocks "github.com/danielblagy/budget-app/internal/db/mocks"
	"github.com/danielblagy/budget-app/internal/model"
	cacheMocks "github.com/danielblagy/budget-app/internal/service/cache/mocks"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Get(t *testing.T) {
	t.Parallel()

	username := "someusername123"
	categoryID := int64(35)
	cacheKey := fmt.Sprintf("%s:category:%d", username, categoryID)

	t.Run("error: can't get category from cache", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		cacheService := new(cacheMocks.Service)
		cacheService.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), cacheKey).
			Return(nil, false, expectedErr)

		service := NewService(nil, nil, cacheService)
		_, err := service.Get(context.Background(), username, categoryID)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't get category from cache")
	})
	t.Run("error: can't unmarshal data", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")
		cacheValueBytes := []byte{1, 2, 3}
		expectedCategory := &model.Category{
			ID:     categoryID,
			UserID: username,
			Name:   "test category name",
			Type:   model.CategoryTypeIncome,
		}

		cacheService := new(cacheMocks.Service)
		cacheService.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), cacheKey).
			Return(cacheValueBytes, true, expectedErr)

		err := json.Unmarshal(cacheValueBytes, expectedCategory)
		require.ErrorContains(t, err, "")
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		expectedCategory := &model.Category{
			ID:     categoryID,
			UserID: username,
			Name:   "test category name",
			Type:   model.CategoryTypeIncome,
		}
		cacheValueBytes, _ := json.Marshal(expectedCategory)
		cacheService := new(cacheMocks.Service)
		cacheService.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), cacheKey).
			Return(cacheValueBytes, true, nil)

		service := NewService(nil, nil, cacheService)
		err := json.Unmarshal(cacheValueBytes, expectedCategory)
		category, _ := service.Get(context.Background(), username, categoryID)
		require.NoError(t, err)
		require.Equal(t, expectedCategory, category)
	})
	t.Run("error: category not found", func(t *testing.T) {
		t.Parallel()

		cacheValueBytes := []byte{1, 2, 3}
		cacheService := new(cacheMocks.Service)
		cacheService.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), cacheKey).
			Return(cacheValueBytes, false, nil)

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
			Return(nil, pgx.ErrNoRows)

		service := NewService(nil, categoriesQuery, cacheService)
		_, err := service.Get(context.Background(), username, categoryID)
		require.ErrorIs(t, err, ErrNotFound)
	})

	t.Run("error: can't get category", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		cacheValueBytes := []byte{1, 2, 3}
		cacheService := new(cacheMocks.Service)
		cacheService.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), cacheKey).
			Return(cacheValueBytes, false, nil)

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
			Return(nil, expectedErr)

		service := NewService(nil, categoriesQuery, cacheService)
		_, err := service.Get(context.Background(), username, categoryID)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't get category")
	})

	t.Run("error: can't set category cache", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		expectedCategory := &model.Category{
			ID:     categoryID,
			UserID: username,
			Name:   "test category name",
			Type:   model.CategoryTypeIncome,
		}

		cacheValueBytes, _ := json.Marshal(expectedCategory)
		cacheService := new(cacheMocks.Service)
		cacheService.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), cacheKey).
			Return(cacheValueBytes, false, nil)

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
			Return(expectedCategory, nil)

		cacheService.
			On("Set", mock.AnythingOfType("*context.emptyCtx"), cacheKey, expectedCategory).
			Return(expectedErr)

		service := NewService(nil, categoriesQuery, cacheService)
		_, err := service.Get(context.Background(), username, categoryID)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't set category cache")
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		expectedCategory := &model.Category{
			ID:     categoryID,
			UserID: username,
			Name:   "test category name",
			Type:   model.CategoryTypeIncome,
		}

		cacheValueBytes, _ := json.Marshal(expectedCategory)
		cacheService := new(cacheMocks.Service)
		cacheService.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), cacheKey).
			Return(cacheValueBytes, false, nil)

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
			Return(expectedCategory, nil)

		cacheService.
			On("Set", mock.AnythingOfType("*context.emptyCtx"), cacheKey, expectedCategory).
			Return(nil)

		service := NewService(nil, categoriesQuery, cacheService)
		category, err := service.Get(context.Background(), username, categoryID)
		require.NoError(t, err)
		require.Equal(t, expectedCategory, category)
	})

}
