package categories

import (
	"context"
	"errors"
	"fmt"
	"testing"

	dbMocks "github.com/danielblagy/budget-app/internal/db/mocks"
	"github.com/danielblagy/budget-app/internal/model"
	cacheMocks "github.com/danielblagy/budget-app/internal/service/cache/mocks"
	"github.com/danielblagy/budget-app/test/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Create(t *testing.T) {
	t.Parallel()

	username := "someusername123"
	category := &model.CreateCategory{
		Name: "some category name",
		Type: model.CategoryTypeExpense,
	}

	t.Run("error: can't create category", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Add", mock.AnythingOfType("*context.emptyCtx"), username, category).
			Return(nil, expectedErr)

		service := NewService(nil, categoriesQuery, nil, nil)
		_, err := service.Create(context.Background(), username, category)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't create category")
	})

	t.Run("error: can't set cache", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		expectedCreatedCategory := &model.Category{
			ID:     10,
			UserID: username,
			Name:   category.Name,
			Type:   category.Type,
		}

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Add", mock.AnythingOfType("*context.emptyCtx"), username, category).
			Return(expectedCreatedCategory, nil)

		cacheService := new(cacheMocks.Service)
		cacheService.
			On("Set", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("%s:category:10", username), expectedCreatedCategory).
			Return(expectedErr)

		logger := new(mocks.Logger)
		logger.On("Error", "can't add to created category to cache", "err", "some error").Once()

		service := NewService(logger, categoriesQuery, cacheService, nil)
		createdCategory, err := service.Create(context.Background(), username, category)
		require.NoError(t, err)
		require.Equal(t, expectedCreatedCategory, createdCategory)
		logger.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		expectedCreatedCategory := &model.Category{
			ID:     10,
			UserID: username,
			Name:   category.Name,
			Type:   category.Type,
		}

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Add", mock.AnythingOfType("*context.emptyCtx"), username, category).
			Return(expectedCreatedCategory, nil)

		cacheService := new(cacheMocks.Service)
		cacheService.
			On("Set", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("%s:category:10", username), expectedCreatedCategory).
			Return(nil)

		service := NewService(nil, categoriesQuery, cacheService, nil)
		createdCategory, err := service.Create(context.Background(), username, category)

		cacheService.AssertExpectations(t)

		require.NoError(t, err)
		require.Equal(t, expectedCreatedCategory, createdCategory)
	})
}
