package db

import (
	"github.com/georgysavva/scany/v2/pgxscan"
)

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=QueryFactory --case=underscore

type QueryFactory interface {
	NewCategoriesQuery(db pgxscan.Querier) CategoriesQuery
	NewEntriesQuery(db pgxscan.Querier) EntriesQuery
}

type queryFactory struct {
}

func NewQueryFactory() QueryFactory {
	return &queryFactory{}
}

func (f *queryFactory) NewCategoriesQuery(db pgxscan.Querier) CategoriesQuery {
	return newCategoriesQuery(db)
}

func (f *queryFactory) NewEntriesQuery(db pgxscan.Querier) EntriesQuery {
	return newEntriesQuery(db)
}
