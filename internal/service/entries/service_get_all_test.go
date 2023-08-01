package entries

import (
	"context"
	"testing"
	"time"

	dbMocks "github.com/danielblagy/budget-app/internal/db/mocks"
	"github.com/danielblagy/budget-app/internal/model"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_GetAll(t *testing.T) {
	t.Parallel()

	username := "someusername"
	entryType := model.EntryTypeExpense

	t.Run("error: can't get entries", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		entriesQuery := new(dbMocks.EntriesQuery)
		entriesQuery.
			On("GetAll", mock.AnythingOfType("*context.emptyCtx"), username, entryType).
			Return(nil, expectedErr)

		service := NewService(entriesQuery)
		_, err := service.GetAll(context.Background(), username, entryType)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't get entries")
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		categoryID1 := int64(20)
		categoryID2 := int64(22)

		expectedEntry := []*model.Entry{
			{
				ID:          int64(1),
				UserID:      username,
				CategoryID:  &categoryID1,
				Amount:      float64(1000.0),
				Date:        time.Now(),
				Description: "eggs,milk,salt and pepper",
				Type:        model.EntryTypeExpense,
			},
			{
				ID:          int64(2),
				UserID:      username,
				CategoryID:  &categoryID2,
				Amount:      float64(150.0),
				Date:        time.Now(),
				Description: "ice cream",
				Type:        model.EntryTypeExpense,
			},
		}

		entriesQuery := new(dbMocks.EntriesQuery)
		entriesQuery.
			On("GetAll", mock.AnythingOfType("*context.emptyCtx"), username, entryType).
			Return(expectedEntry, nil)

		service := NewService(entriesQuery)
		entry, err := service.GetAll(context.Background(), username, entryType)
		require.NoError(t, err)
		require.Equal(t, expectedEntry, entry)
	})

}
