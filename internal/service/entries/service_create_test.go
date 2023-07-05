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

func Test_Create(t *testing.T) {
	t.Parallel()

	username := "someusername123"
	entry := &model.CreateEntry{
		CategoryID:  int64(20),
		Amount:      float64(1000.0),
		Date:        "2023-07-02",
		Description: "eggs,milk,salt and pepper",
		Type:        model.EntryTypeExpense,
	}

	t.Run("error:can't create entry", func(t *testing.T) {
		t.Parallel()

		expectedErr := errors.New("some error")

		entriesQuery := new(dbMocks.EntriesQuery)
		entriesQuery.
			On("Add", mock.AnythingOfType("*context.emptyCtx"), username, entry).
			Return(nil, expectedErr)

		service := NewService(entriesQuery)
		_, err := service.Create(context.Background(), username, entry)
		require.ErrorIs(t, err, expectedErr)
		require.ErrorContains(t, err, "can't create entry")
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		expectedCreatedEntry := &model.Entry{
			ID:          int64(1),
			UserID:      username,
			CategoryID:  int64(20),
			Amount:      float64(1000.0),
			Date:        time.Now(),
			Description: "eggs,milk,salt and pepper",
			Type:        model.EntryTypeExpense,
		}

		entriesQuery := new(dbMocks.EntriesQuery)
		entriesQuery.
			On("Add", mock.AnythingOfType("*context.emptyCtx"), username, entry).
			Return(expectedCreatedEntry, nil)

		service := NewService(entriesQuery)
		createdEntry, err := service.Create(context.Background(), username, entry)

		require.NoError(t, err)
		require.Equal(t, expectedCreatedEntry, createdEntry)
	})

}
