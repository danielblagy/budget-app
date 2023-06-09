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

func Test_GetAll(t *testing.T) {
	t.Parallel()

	username := "someusername123"
	categoryType := model.CategoryTypeExpense

	t.Run("error: can't get categories", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("GetAll", mock.AnythingOfType("*context.emptyCtx"), username, categoryType).
			Return(nil, expectedErr)

		service := NewService(categoriesQuery)
		_, err := service.GetAll(context.Background(), username, categoryType)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't get categories")
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		expectedCategory := []*model.Category{
			{
				ID:     10,
				UserID: username,
				Name:   "test category name",
				Type:   model.CategoryTypeIncome,
			},
		}

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("GetAll", mock.AnythingOfType("*context.emptyCtx"), username, categoryType).
			Return(expectedCategory, nil)

		service := NewService(categoriesQuery)
		category, err := service.GetAll(context.Background(), username, categoryType)
		require.NoError(t, err)
		require.Equal(t, expectedCategory, category)
	})

}
