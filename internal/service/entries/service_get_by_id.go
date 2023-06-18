package entries

import (
	"context"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("entry not found")

func (s service) GetByID(ctx context.Context, username string, entryID int64) (*model.Entry, error) {
	entry, err := s.entriesQuery.GetByID(ctx, username, entryID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, errors.Wrap(err, "can't get entry")
	}
	return entry, nil
}
