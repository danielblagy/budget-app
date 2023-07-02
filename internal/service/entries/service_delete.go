package entries

import (
	"context"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (s service) Delete(ctx context.Context, username string, entryID int64) (*model.Entry, error) {
	deletedEntry, err := s.entriesQuery.Delete(ctx, username, entryID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, errors.Wrap(err, "can't delete entry")
	}
	return deletedEntry, nil
}
