package entries

import (
	"context"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/pkg/errors"
)

func (s service) GetAll(ctx context.Context, username string, entryType model.EntryType) ([]*model.Entry, error) {
	entries, err := s.entriesQuery.GetAll(ctx, username, entryType)
	if err != nil {
		return nil, errors.Wrap(err, "can't get entries")
	}
	return entries, nil

}
