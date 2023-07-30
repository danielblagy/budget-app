package categories

import (
	"context"
	"fmt"
	"testing"

	dbMocks "github.com/danielblagy/budget-app/internal/db/mocks"
	"github.com/danielblagy/budget-app/internal/model"
	cacheMocks "github.com/danielblagy/budget-app/internal/service/cache/mocks"
	"github.com/danielblagy/budget-app/test/mocks"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Delete(t *testing.T) {
	t.Parallel()

	username := "someusername123"
	categoryID := int64(35)

	t.Run("error: can't begin tx", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		dbRunner := new(dbMocks.DbRunner)
		dbRunner.On("Begin", mock.AnythingOfType("*context.emptyCtx")).Return(nil, expectedErr)

		queryFactory := new(dbMocks.QueryFactory)
		queryFactory.On("GetDbRunner").Return(dbRunner)

		service := NewService(nil, nil, nil, queryFactory)
		_, err := service.Delete(context.Background(), username, categoryID, false)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't begin tx")
	})

	t.Run("error: category not found", func(t *testing.T) {
		t.Parallel()

		tx := new(mocks.Tx)
		tx.On("Rollback", mock.AnythingOfType("*context.emptyCtx")).Return(nil)

		dbRunner := new(dbMocks.DbRunner)
		dbRunner.On("Begin", mock.AnythingOfType("*context.emptyCtx")).Return(tx, nil)

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Delete", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
			Return(nil, pgx.ErrNoRows)

		queryFactory := new(dbMocks.QueryFactory)
		queryFactory.On("GetDbRunner").Return(dbRunner)
		queryFactory.On("NewCategoriesQuery", tx).Return(categoriesQuery)

		service := NewService(nil, nil, nil, queryFactory)
		_, err := service.Delete(context.Background(), username, categoryID, false)
		require.ErrorIs(t, err, ErrNotFound)
	})

	t.Run("error: can't delete category", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		tx := new(mocks.Tx)
		tx.On("Rollback", mock.AnythingOfType("*context.emptyCtx")).Return(nil)

		dbRunner := new(dbMocks.DbRunner)
		dbRunner.On("Begin", mock.AnythingOfType("*context.emptyCtx")).Return(tx, nil)

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Delete", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
			Return(nil, expectedErr)

		queryFactory := new(dbMocks.QueryFactory)
		queryFactory.On("GetDbRunner").Return(dbRunner)
		queryFactory.On("NewCategoriesQuery", tx).Return(categoriesQuery)

		service := NewService(nil, nil, nil, queryFactory)
		_, err := service.Delete(context.Background(), username, categoryID, false)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't delete category")
	})

	t.Run("error: can't delete category from cache", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		tx := new(mocks.Tx)
		tx.On("Rollback", mock.AnythingOfType("*context.emptyCtx")).Return(nil)

		dbRunner := new(dbMocks.DbRunner)
		dbRunner.On("Begin", mock.AnythingOfType("*context.emptyCtx")).Return(tx, nil)

		categoriesQuery := new(dbMocks.CategoriesQuery)
		categoriesQuery.
			On("Delete", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
			Return(&model.Category{
				ID:     categoryID,
				UserID: username,
				Name:   "test category name",
				Type:   model.CategoryTypeIncome,
			}, nil)

		queryFactory := new(dbMocks.QueryFactory)
		queryFactory.On("GetDbRunner").Return(dbRunner)
		queryFactory.On("NewCategoriesQuery", tx).Return(categoriesQuery)

		cacheService := new(cacheMocks.Service)
		cacheService.
			On("Delete", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("%s:category:%d", username, categoryID)).
			Return(expectedErr)

		service := NewService(nil, nil, cacheService, queryFactory)
		_, err := service.Delete(context.Background(), username, categoryID, false)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't delete category from cache")
	})

	t.Run("delete entries", func(t *testing.T) {
		t.Parallel()

		t.Run("error: can't delete category entries", func(t *testing.T) {
			t.Parallel()

			expectedErr := errors.New("some error")

			tx := new(mocks.Tx)
			tx.On("Rollback", mock.AnythingOfType("*context.emptyCtx")).Return(nil)

			dbRunner := new(dbMocks.DbRunner)
			dbRunner.On("Begin", mock.AnythingOfType("*context.emptyCtx")).Return(tx, nil)

			categoriesQuery := new(dbMocks.CategoriesQuery)
			categoriesQuery.
				On("Delete", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
				Return(&model.Category{
					ID:     categoryID,
					UserID: username,
					Name:   "test category name",
					Type:   model.CategoryTypeIncome,
				}, nil)

			entriesQuery := new(dbMocks.EntriesQuery)
			entriesQuery.
				On("DeleteByUserAndCategory", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
				Return(expectedErr)

			queryFactory := new(dbMocks.QueryFactory)
			queryFactory.On("GetDbRunner").Return(dbRunner)
			queryFactory.On("NewCategoriesQuery", tx).Return(categoriesQuery)
			queryFactory.On("NewEntriesQuery", tx).Return(entriesQuery)

			cacheService := new(cacheMocks.Service)
			cacheService.
				On("Delete", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("%s:category:%d", username, categoryID)).
				Return(nil)

			service := NewService(nil, nil, cacheService, queryFactory)
			_, err := service.Delete(context.Background(), username, categoryID, true)
			require.ErrorIs(t, err, expectedErr)
			require.ErrorContains(t, err, "can't delete category entries")
		})

		t.Run("success", func(t *testing.T) {
			t.Parallel()

			expectedCategory := &model.Category{
				ID:     categoryID,
				UserID: username,
				Name:   "test category name",
				Type:   model.CategoryTypeIncome,
			}

			tx := new(mocks.Tx)
			tx.On("Commit", mock.AnythingOfType("*context.emptyCtx")).Return(nil)

			dbRunner := new(dbMocks.DbRunner)
			dbRunner.On("Begin", mock.AnythingOfType("*context.emptyCtx")).Return(tx, nil)

			categoriesQuery := new(dbMocks.CategoriesQuery)
			categoriesQuery.
				On("Delete", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
				Return(expectedCategory, nil)

			entriesQuery := new(dbMocks.EntriesQuery)
			entriesQuery.
				On("DeleteByUserAndCategory", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
				Return(nil)

			queryFactory := new(dbMocks.QueryFactory)
			queryFactory.On("GetDbRunner").Return(dbRunner)
			queryFactory.On("NewCategoriesQuery", tx).Return(categoriesQuery)
			queryFactory.On("NewEntriesQuery", tx).Return(entriesQuery)

			cacheService := new(cacheMocks.Service)
			cacheService.
				On("Delete", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("%s:category:%d", username, categoryID)).
				Return(nil)

			service := NewService(nil, nil, cacheService, queryFactory)
			category, err := service.Delete(context.Background(), username, categoryID, true)
			require.NoError(t, err)
			require.Equal(t, expectedCategory, category)
		})
	})

	t.Run("set enrties category to null", func(t *testing.T) {
		t.Parallel()

		t.Run("error: can't set category entries's category_id to null", func(t *testing.T) {
			t.Parallel()

			expectedErr := errors.New("some error")

			tx := new(mocks.Tx)
			tx.On("Rollback", mock.AnythingOfType("*context.emptyCtx")).Return(nil)

			dbRunner := new(dbMocks.DbRunner)
			dbRunner.On("Begin", mock.AnythingOfType("*context.emptyCtx")).Return(tx, nil)

			categoriesQuery := new(dbMocks.CategoriesQuery)
			categoriesQuery.
				On("Delete", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
				Return(&model.Category{
					ID:     categoryID,
					UserID: username,
					Name:   "test category name",
					Type:   model.CategoryTypeIncome,
				}, nil)

			entriesQuery := new(dbMocks.EntriesQuery)
			entriesQuery.
				On("SetNullCategory", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
				Return(expectedErr)

			queryFactory := new(dbMocks.QueryFactory)
			queryFactory.On("GetDbRunner").Return(dbRunner)
			queryFactory.On("NewCategoriesQuery", tx).Return(categoriesQuery)
			queryFactory.On("NewEntriesQuery", tx).Return(entriesQuery)

			cacheService := new(cacheMocks.Service)
			cacheService.
				On("Delete", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("%s:category:%d", username, categoryID)).
				Return(nil)

			service := NewService(nil, nil, cacheService, queryFactory)
			_, err := service.Delete(context.Background(), username, categoryID, false)
			require.ErrorIs(t, err, expectedErr)
			require.ErrorContains(t, err, "can't set category entries's category_id to null")
		})

		t.Run("error: can't commit tx", func(t *testing.T) {
			t.Parallel()

			expectedErr := errors.New("some error")

			tx := new(mocks.Tx)
			tx.On("Commit", mock.AnythingOfType("*context.emptyCtx")).Return(expectedErr)

			dbRunner := new(dbMocks.DbRunner)
			dbRunner.On("Begin", mock.AnythingOfType("*context.emptyCtx")).Return(tx, nil)

			categoriesQuery := new(dbMocks.CategoriesQuery)
			categoriesQuery.
				On("Delete", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
				Return(&model.Category{
					ID:     categoryID,
					UserID: username,
					Name:   "test category name",
					Type:   model.CategoryTypeIncome,
				}, nil)

			entriesQuery := new(dbMocks.EntriesQuery)
			entriesQuery.
				On("SetNullCategory", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
				Return(nil)

			queryFactory := new(dbMocks.QueryFactory)
			queryFactory.On("GetDbRunner").Return(dbRunner)
			queryFactory.On("NewCategoriesQuery", tx).Return(categoriesQuery)
			queryFactory.On("NewEntriesQuery", tx).Return(entriesQuery)

			cacheService := new(cacheMocks.Service)
			cacheService.
				On("Delete", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("%s:category:%d", username, categoryID)).
				Return(nil)

			service := NewService(nil, nil, cacheService, queryFactory)
			_, err := service.Delete(context.Background(), username, categoryID, false)
			require.ErrorIs(t, err, expectedErr)
			require.ErrorContains(t, err, "can't commit tx")
		})

		t.Run("success", func(t *testing.T) {
			t.Parallel()

			expectedCategory := &model.Category{
				ID:     categoryID,
				UserID: username,
				Name:   "test category name",
				Type:   model.CategoryTypeIncome,
			}

			tx := new(mocks.Tx)
			tx.On("Commit", mock.AnythingOfType("*context.emptyCtx")).Return(nil)

			dbRunner := new(dbMocks.DbRunner)
			dbRunner.On("Begin", mock.AnythingOfType("*context.emptyCtx")).Return(tx, nil)

			categoriesQuery := new(dbMocks.CategoriesQuery)
			categoriesQuery.
				On("Delete", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
				Return(expectedCategory, nil)

			entriesQuery := new(dbMocks.EntriesQuery)
			entriesQuery.
				On("SetNullCategory", mock.AnythingOfType("*context.emptyCtx"), username, categoryID).
				Return(nil)

			queryFactory := new(dbMocks.QueryFactory)
			queryFactory.On("GetDbRunner").Return(dbRunner)
			queryFactory.On("NewCategoriesQuery", tx).Return(categoriesQuery)
			queryFactory.On("NewEntriesQuery", tx).Return(entriesQuery)

			cacheService := new(cacheMocks.Service)
			cacheService.
				On("Delete", mock.AnythingOfType("*context.emptyCtx"), fmt.Sprintf("%s:category:%d", username, categoryID)).
				Return(nil)

			service := NewService(nil, nil, cacheService, queryFactory)
			category, err := service.Delete(context.Background(), username, categoryID, false)
			require.NoError(t, err)
			require.Equal(t, expectedCategory, category)
		})
	})
}
