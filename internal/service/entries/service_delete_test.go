package entries

import (
	"context"
	"testing"
	"time"

	dbMocks "github.com/danielblagy/budget-app/internal/db/mocks"
	"github.com/danielblagy/budget-app/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Delete(t *testing.T) {
	t.Parallel()

	username := "someusername123"
	entryID := int64(1)

	t.Run("error: entry not found", func(t *testing.T) {
		t.Parallel()

		entriesQuery := new(dbMocks.EntriesQuery)
		entriesQuery.
			On("Delete", mock.AnythingOfType("*context.emptyCtx"), username, entryID).
			Return(nil, pgx.ErrNoRows)

		service := NewService(entriesQuery)
		_, err := service.Delete(context.Background(), username, entryID)
		require.ErrorIs(t, err, ErrNotFound)
	})

	t.Run("error: can't delete entry", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		entriesQuery := new(dbMocks.EntriesQuery)
		entriesQuery.
			On("Delete", mock.AnythingOfType("*context.emptyCtx"), username, entryID).
			Return(nil, expectedErr)

		service := NewService(entriesQuery)
		_, err := service.Delete(context.Background(), username, entryID)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't delete entry")
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		expectedEntry := &model.Entry{
			ID:          entryID,
			UserID:      username,
			CategoryID:  int64(20),
			Amount:      float64(1000.0),
			Date:        time.Now(),
			Description: "eggs,milk,salt and pepper",
			Type:        model.EntryTypeExpense,
		}

		entriesQuery := new(dbMocks.EntriesQuery)
		entriesQuery.
			On("Delete", mock.AnythingOfType("*context.emptyCtx"), username, entryID).
			Return(expectedEntry, nil)

		service := NewService(entriesQuery)
		entry, err := service.Delete(context.Background(), username, entryID)
		require.NoError(t, err)
		require.Equal(t, expectedEntry, entry)
	})

}
