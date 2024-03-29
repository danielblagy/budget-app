package entries

import (
	"context"

	"github.com/danielblagy/budget-app/internal/db"
	"github.com/danielblagy/budget-app/internal/model"
)

type Service interface {
	// GetAll returns all users's entries of type.
	GetAll(ctx context.Context, username string, entryType model.EntryType) ([]*model.Entry, error)
	GetByID(ctx context.Context, username string, entryID int64) (*model.Entry, error)
	Create(ctx context.Context, username string, entry *model.CreateEntry) (*model.Entry, error)
	Update(ctx context.Context, username string, entryID int64, entry *model.UpdateEntry) (*model.Entry, error)
	Delete(ctx context.Context, username string, entryID int64) (*model.Entry, error)
}

type service struct {
	entriesQuery db.EntriesQuery
}

func NewService(entriesQuery db.EntriesQuery) Service {
	return &service{
		entriesQuery: entriesQuery,
	}
}
