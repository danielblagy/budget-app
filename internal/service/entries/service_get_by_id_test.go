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

func Test_GetById(t *testing.T) {
	t.Parallel()

	username := "someusername123"
	entryID := int64(1)

	t.Run("error:can't get entry", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		entriesQuery := new(dbMocks.EntriesQuery)
		entriesQuery.
			On("GetByID", mock.AnythingOfType("*context.emptyCtx"), username, entryID).
			Return(nil, expectedErr)

		service := NewService(entriesQuery)
		_, err := service.GetByID(context.Background(), username, entryID)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't get entry")
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		expectedEntry := &model.Entry{
			ID:          entryID,
			UserID:      username,
			CategoryID:  int64(20),
			Amount:      float64(1000.0),
			Date:        time.Time(time.Now()),
			Description: "eggs,milk,salt and pepper",
			Type:        model.EntryTypeExpense,
		}

		entriesQuery := new(dbMocks.EntriesQuery)
		entriesQuery.
			On("GetByID", mock.AnythingOfType("*context.emptyCtx"), username, entryID).
			Return(expectedEntry, nil)

		service := NewService(entriesQuery)
		entry, err := service.GetByID(context.Background(), username, entryID)
		require.NoError(t, err)
		require.Equal(t, expectedEntry, entry)
	})

}
