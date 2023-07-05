package entries

import (
	"context"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (s service) Update(ctx context.Context, username string, entryID int64, updateData *model.UpdateEntry) (*model.Entry, error) {
	updatedEntry, err := s.entriesQuery.Update(ctx, username, entryID, updateData.CategoryID, updateData.Amount, updateData.Date, updateData.Description, updateData.Type)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, errors.Wrap(err, "can't update entry")
	}

	return updatedEntry, nil
}
