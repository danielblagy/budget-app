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

func Test_Update(t *testing.T) {
	t.Parallel()

	username := "someusername123"
	category := &model.UpdateCategory{
		ID:   int64(12),
		Name: "some updated category name",
	}

	t.Run("error: category not found", func(t *testing.T) {
		t.Parallel()

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Update", mock.AnythingOfType("*context.emptyCtx"), username, category.ID, category.Name).
			Return(nil, pgx.ErrNoRows)

		service := NewService(nil, categoriesQuery, nil)
		_, err := service.Update(context.Background(), username, category)
		require.ErrorIs(t, err, ErrNotFound)
	})

	t.Run("error: can't update category", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Update", mock.AnythingOfType("*context.emptyCtx"), username, category.ID, category.Name).
			Return(nil, expectedErr)

		service := NewService(nil, categoriesQuery, nil)
		_, err := service.Update(context.Background(), username, category)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't update category")
	})

	t.Run("error: can't update category cache", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		expectedUpdatedCategory := &model.Category{
			ID:     category.ID,
			UserID: username,
			Name:   category.Name,
			Type:   model.CategoryTypeIncome,
		}

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Update", mock.AnythingOfType("*context.emptyCtx"), username, category.ID, category.Name).
			Return(expectedUpdatedCategory, nil)

		cacheService := new(cacheMocks.Service)
		cacheService.
			On("Set", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("%s:category:%d", username, category.ID), expectedUpdatedCategory).
			Return(expectedErr)

		service := NewService(nil, categoriesQuery, cacheService)
		_, err := service.Update(context.Background(), username, category)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't update category cache")
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		expectedUpdatedCategory := &model.Category{
			ID:     category.ID,
			UserID: username,
			Name:   category.Name,
			Type:   model.CategoryTypeIncome,
		}

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Update", mock.AnythingOfType("*context.emptyCtx"), username, category.ID, category.Name).
			Return(expectedUpdatedCategory, nil)

		cacheService := new(cacheMocks.Service)
		cacheService.
			On("Set", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("%s:category:%d", username, category.ID), expectedUpdatedCategory).
			Return(nil)

		service := NewService(nil, categoriesQuery, cacheService)
		updatedCategory, err := service.Update(context.Background(), username, category)
		require.NoError(t, err)
		require.Equal(t, expectedUpdatedCategory, updatedCategory)
	})
}
