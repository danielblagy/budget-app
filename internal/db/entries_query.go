package db

import (
	"context"
	"fmt"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=EntriesQuery --case=underscore

type EntriesQuery interface {
	GetAll(ctx context.Context, username string, entryType model.EntryType) ([]*model.Entry, error)
	GetByID(ctx context.Context, username string, entryID int64) (*model.Entry, error)
}
type entriesQuery struct {
	db *pgx.Conn
}

func NewEntriesQuery(db *pgx.Conn) EntriesQuery {
	return &entriesQuery{
		db: db,
	}
}

func (q entriesQuery) GetAll(ctx context.Context, username string, entryType model.EntryType) ([]*model.Entry, error) {
	var getAllQueryTemplate = "select id, user_id, category_id, amount, date, description, type from entries where user_id = '%s' and type = '%s'"

	var entries []*model.Entry
	err := pgxscan.Select(ctx, q.db, &entries, fmt.Sprintf(getAllQueryTemplate, username, entryType))
	if err != nil {
		return nil, errors.Wrap(err, "can't get entries from db")
	}

	return entries, nil
}

func (q entriesQuery) GetByID(ctx context.Context, username string, entryID int64) (*model.Entry, error) {
	var ErrNotFound = errors.New("category not found")
	var getQueryTemplate = "select id, user_id, name, type from categories where id = '%d' and user_id = '%s'"

	var entry model.Entry

	err := pgxscan.Get(ctx, q.db, &entry, fmt.Sprintf(getQueryTemplate, entryID, username))
	if err != nil {
		return nil, ErrNotFound
	}

	return &entry, nil
}
