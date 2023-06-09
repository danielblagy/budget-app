package categories

import (
	"context"
	"testing"

	dbMocks "github.com/danielblagy/budget-app/internal/db/mocks"
	"github.com/danielblagy/budget-app/internal/model"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Delete(t *testing.T) {
	t.Parallel()

	username := "someusername123"
	categoryID := int64(35)

	t.Run("error: can't delete category", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Delete", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
			Return(nil, expectedErr)

		service := NewService(categoriesQuery, nil)
		_, err := service.Delete(context.Background(), username, categoryID)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't delete category")
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
			On("Delete", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
			Return(expectedCategory, nil)

		service := NewService(categoriesQuery, nil)
		category, err := service.Delete(context.Background(), username, categoryID)
		require.NoError(t, err)
		require.Equal(t, expectedCategory, category)
	})
}
