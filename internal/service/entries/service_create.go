package entries

import (
	"context"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/pkg/errors"
)

func (s service) Create(ctx context.Context, username string, entry *model.CreateEntry) (*model.Entry, error) {
	createdEntry, err := s.entriesQuery.Add(ctx, username, entry)

	if err != nil {
		return nil, errors.Wrap(err, "can't create entry")
	}

	return createdEntry, nil
}
