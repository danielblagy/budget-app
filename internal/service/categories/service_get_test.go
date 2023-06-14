package categories

import (
	"context"
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

	t.Run("error: category not found", func(t *testing.T) {
		t.Parallel()

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
			Return(nil, pgx.ErrNoRows)

		service := NewService(nil, categoriesQuery, nil)
		_, err := service.Get(context.Background(), username, categoryID)
		require.ErrorIs(t, err, ErrNotFound)
	})

	t.Run("error: can't get category", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
			Return(nil, expectedErr)

		service := NewService(nil, categoriesQuery, nil)
		_, err := service.Get(context.Background(), username, categoryID)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't get category")
	})

	t.Run("error: can't get category from cache", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		expectedCategory := &model.Category{
			ID:     categoryID,
			UserID: username,
			Name:   "test category name",
			Type:   model.CategoryTypeIncome,
		}

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
			Return(&model.Category{
				ID:     categoryID,
				UserID: username,
				Name:   "test category name",
				Type:   model.CategoryTypeIncome,
			}, nil)

		cacheService := new(cacheMocks.Service)
		cacheService.
			On("Set", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("%s:category:%d", username, categoryID), expectedCategory).
			Return(expectedErr)

		service := NewService(nil, categoriesQuery, cacheService)
		_, err := service.Get(context.Background(), username, categoryID)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't get category from cache")
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		expectedCategory := &model.Category{
			ID:     categoryID,
			UserID: username,
			Name:   "test category name",
			Type:   model.CategoryTypeIncome,
		}

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
			Return(expectedCategory, nil)

		cacheService := new(cacheMocks.Service)
		cacheService.
			On("Set", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("%s:category:%d", username, categoryID), expectedCategory).
			Return(nil)

		service := NewService(nil, categoriesQuery, cacheService)
		category, err := service.Get(context.Background(), username, categoryID)
		require.NoError(t, err)
		require.Equal(t, expectedCategory, category)
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		expectedCategory := &model.Category{
			ID:     categoryID,
			UserID: username,
			Name:   "test category name",
			Type:   model.CategoryTypeIncome,
		}

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
			Return(expectedCategory, nil)

		service := NewService(nil, categoriesQuery, nil)
		category, err := service.Get(context.Background(), username, categoryID)
		require.NoError(t, err)
		require.Equal(t, expectedCategory, category)
	})
}
