package entries

import (
	"context"
	"fmt"
	"testing"
	"time"

	dbMocks "github.com/danielblagy/budget-app/internal/db/mocks"
	"github.com/danielblagy/budget-app/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func parseTime(s string) (time.Time, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, s)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse time: %v", err)
	}
	return t, nil
}

func Test_Update(t *testing.T) {
	t.Parallel()

	username := "someusername123"
	entryID := int64(5)
	entry := &model.UpdateEntry{
		CategoryID:  int64(20),
		Amount:      float64(1000.0),
		Date:        "2023-07-02",
		Description: "eggs,milk,salt and pepper",
		Type:        model.EntryTypeExpense,
	}

	t.Run("error:entry not found", func(t *testing.T) {
		t.Parallel()

		entriesQuery := new(dbMocks.EntriesQuery)
		entriesQuery.
			On("Update", mock.AnythingOfType("*context.emptyCtx"), username, entryID, entry.CategoryID, entry.Amount, entry.Date, entry.Description, entry.Type).
			Return(nil, pgx.ErrNoRows)

		service := NewService(entriesQuery)
		_, err := service.Update(context.Background(), username, entryID, entry)
		require.ErrorIs(t, err, ErrNotFound)
	})

	t.Run("error: can't update entry", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		entriesQuery := new(dbMocks.EntriesQuery)
		entriesQuery.
			On("Update", mock.AnythingOfType("*context.emptyCtx"), username, entryID, entry.CategoryID, entry.Amount, entry.Date, entry.Description, entry.Type).
			Return(nil, expectedErr)

		service := NewService(entriesQuery)
		_, err := service.Update(context.Background(), username, entryID, entry)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't update entry")
	})

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		expectedDate, _ := parseTime(entry.Date)

		expectedUpdatedEntry := &model.Entry{
			ID:          entryID,
			UserID:      username,
			CategoryID:  entry.CategoryID,
			Amount:      entry.Amount,
			Date:        expectedDate,
			Description: entry.Description,
			Type:        entry.Type,
		}

		entriesQuery := new(dbMocks.EntriesQuery)
		entriesQuery.
			On("Update", mock.AnythingOfType("*context.emptyCtx"), username, entryID, entry.CategoryID, entry.Amount, entry.Date, entry.Description, entry.Type).
			Return(expectedUpdatedEntry, nil)

		service := NewService(entriesQuery)
		updatedEntry, err := service.Update(context.Background(), username, entryID, entry)
		require.NoError(t, err)
		require.Equal(t, expectedUpdatedEntry, updatedEntry)

	})
}
