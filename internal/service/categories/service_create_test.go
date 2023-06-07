package categories

import (
	"context"
	"errors"
	"testing"

	dbMocks "github.com/danielblagy/budget-app/internal/db/mocks"
	"github.com/danielblagy/budget-app/internal/model"
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

		service := NewService(categoriesQuery)
		_, err := service.Create(context.Background(), username, category)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't create category")
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

		service := NewService(categoriesQuery)
		createdCategory, err := service.Create(context.Background(), username, category)
		require.Equal(t, expectedCreatedCategory, createdCategory)
		require.NoError(t, err)
	})
}
