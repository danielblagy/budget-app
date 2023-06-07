package db

import (
	"context"
	"fmt"

	"github.com/danielblagy/budget-app/internal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=CategoriesQuery --case=underscore

type CategoriesQuery interface {
	GetAll(ctx context.Context, username string, categoryType model.CategoryType) ([]*model.Category, error)
	Get(ctx context.Context, username string, categoryID int64) (*model.Category, error)
	Add(ctx context.Context, username string, category *model.CreateCategory) (*model.Category, error)
	Update(ctx context.Context, username string, categoryID int64, newName string) (*model.Category, error)
	Delete(ctx context.Context, username string, categoryID int64) (*model.Category, error)
}

type categoriesQuery struct {
	db *pgx.Conn
}

func NewCategoriesQuery(db *pgx.Conn) CategoriesQuery {
	return &categoriesQuery{
		db: db,
	}
}

func (q categoriesQuery) GetAll(ctx context.Context, username string, categoryType model.CategoryType) ([]*model.Category, error) {
	var getAllQueryTemplate = "select id, user_id, name, type from categories where user_id = '%s' and type = '%s'"

	var categories []*model.Category
	err := pgxscan.Select(ctx, q.db, &categories, fmt.Sprintf(getAllQueryTemplate, username, categoryType))
	if err != nil {
		return nil, errors.Wrap(err, "can't get categories from db")
	}

	return categories, nil
}

func (q categoriesQuery) Get(ctx context.Context, username string, categoryID int64) (*model.Category, error) {
	var getQueryTemplate = "select id, user_id, name, type from categories where id = '%d' and user_id = '%s'"

	var category model.Category
	err := pgxscan.Get(ctx, q.db, &category, fmt.Sprintf(getQueryTemplate, categoryID, username))
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (q categoriesQuery) Add(ctx context.Context, username string, category *model.CreateCategory) (*model.Category, error) {
	var addQueryTemplate = "insert into categories (user_id, name, type) values ('%s', '%s', '%s') returning id, user_id, name, type"

	rows, err := q.db.Query(ctx, fmt.Sprintf(addQueryTemplate, username, category.Name, category.Type))
	if err != nil {
		return nil, err
	}

	var createdCategory model.Category
	err = pgxscan.ScanOne(&createdCategory, rows)
	if err != nil {
		return nil, err
	}

	return &createdCategory, nil
}

func (q categoriesQuery) Update(ctx context.Context, username string, categoryID int64, newName string) (*model.Category, error) {
	var updateQueryTemplate = "update categories set name = '%s' where id = '%d' and user_id = '%s' returning id, user_id, name, type"

	var updatedCategory model.Category
	err := pgxscan.Get(ctx, q.db, &updatedCategory, fmt.Sprintf(updateQueryTemplate, newName, categoryID, username))
	if err != nil {
		return nil, err
	}

	return &updatedCategory, nil
}

func (q categoriesQuery) Delete(ctx context.Context, username string, categoryID int64) (*model.Category, error) {
	var deleteQueryTemplate = "delete from categories where user_id = '%s' and id = '%d' returning id, user_id, name, type"

	var deletedCategory model.Category
	err := pgxscan.Get(ctx, q.db, &deletedCategory, fmt.Sprintf(deleteQueryTemplate, username, categoryID))
	if err != nil {
		return nil, err
	}

	return &deletedCategory, nil
}
