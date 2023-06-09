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

func Test_Update(t *testing.T) {
	t.Parallel()

	username := "someusername123"
	category := &model.UpdateCategory{
		ID:   int64(12),
		Name: "some updated category name",
	}

	t.Run("error: can't update categories", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Update", mock.AnythingOfType("*context.emptyCtx"), username, category.ID, category.Name).
			Return(nil, expectedErr)

		service := NewService(categoriesQuery, nil)
		_, err := service.Update(context.Background(), username, category)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't update category")
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

		service := NewService(categoriesQuery, nil)
		updatedCategory, err := service.Update(context.Background(), username, category)
		require.NoError(t, err)
		require.Equal(t, expectedUpdatedCategory, updatedCategory)
	})
}
