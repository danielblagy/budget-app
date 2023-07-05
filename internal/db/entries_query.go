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
	Add(ctx context.Context, username string, entry *model.CreateEntry) (*model.Entry, error)
	Update(ctx context.Context, username string, entryID int64, newCategoryID int64, newAmount float64, newDate string, newDescription string, newType model.EntryType) (*model.Entry, error)
	Delete(ctx context.Context, username string, entryID int64) (*model.Entry, error)
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
	var getQueryTemplate = "select id, user_id, category_id, amount, date, description, type from entries where id = '%d' and user_id = '%s'"

	var entry model.Entry

	err := pgxscan.Get(ctx, q.db, &entry, fmt.Sprintf(getQueryTemplate, entryID, username))
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func (q entriesQuery) Add(ctx context.Context, username string, entry *model.CreateEntry) (*model.Entry, error) {
	var addQueryTemplate = "insert into entries (user_id, category_id, amount, date, description, type) values ('%s','%d','%f','%s','%s','%s') returning id,user_id, category_id, amount, date, description, type"

	rows, err := q.db.Query(ctx, fmt.Sprintf(addQueryTemplate, username, entry.CategoryID, entry.Amount, entry.Date, entry.Description, entry.Type))
	if err != nil {
		return nil, err
	}

	var createdEntry model.Entry
	err = pgxscan.ScanOne(&createdEntry, rows)
	if err != nil {
		return nil, err
	}

	return &createdEntry, nil
}

func (q entriesQuery) Update(ctx context.Context, username string, entryID int64, newCategoryID int64, newAmount float64, newDate string, newDescription string, newType model.EntryType) (*model.Entry, error) {
	var updateQueryTemplate = "update entries set category_id ='%d',amount='%f',date='%s',description = '%s',type = '%s' where user_id='%s' and id='%d' returning id,user_id,category_id,amount,date,description,type"

	var updatedEntry model.Entry
	err := pgxscan.Get(ctx, q.db, &updatedEntry, fmt.Sprintf(updateQueryTemplate, newCategoryID, newAmount, newDate, newDescription, newType, username, entryID))

	if err != nil {
		return nil, err
	}

	return &updatedEntry, nil
}

func (q entriesQuery) Delete(ctx context.Context, username string, entryID int64) (*model.Entry, error) {
	var deleteQueryTemplate = "delete from entries where user_id = '%s' and id = '%d' returning id, user_id, category_id, amount, date, description, type"

	var deletedEntry model.Entry
	err := pgxscan.Get(ctx, q.db, &deletedEntry, fmt.Sprintf(deleteQueryTemplate, username, entryID))
	if err != nil {
		return nil, err
	}

	return &deletedEntry, nil
}
