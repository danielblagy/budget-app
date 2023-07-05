package categories

import (
	"context"
	"testing"

	dbMocks "github.com/danielblagy/budget-app/internal/db/mocks"
	"github.com/danielblagy/budget-app/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Exists(t *testing.T) {
	t.Parallel()

	username := "someusername123"
	categoryID := int64(5)

	t.Run("error: can't check if category exists", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
			Return(nil, expectedErr)

		service := NewService(nil, categoriesQuery, nil)
		_, err := service.Exists(context.Background(), username, categoryID)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't check if category exists")
	})
	t.Run("error:category not found", func(t *testing.T) {
		t.Parallel()

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
			Return(nil, pgx.ErrNoRows)

		service := NewService(nil, categoriesQuery, nil)
		_, err := service.Exists(context.Background(), username, categoryID)
		require.ErrorIs(t, err, nil)
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

		var exists bool = true
		service := NewService(nil, categoriesQuery, nil)
		isExist, err := service.Exists(context.Background(), username, categoryID)
		require.NoError(t, err)
		require.Equal(t, exists, isExist)
	})
	t.Run("success", func(t *testing.T) {
		t.Parallel()

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Get", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
			Return(nil, pgx.ErrNoRows)

		var exists bool = false
		service := NewService(nil, categoriesQuery, nil)
		isExist, err := service.Exists(context.Background(), username, categoryID)
		require.NoError(t, err)
		require.Equal(t, exists, isExist)

	})
}
