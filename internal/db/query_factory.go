package db

import (
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=QueryFactory --case=underscore

type QueryFactory interface {
	GetDbRunner() DbRunner
	NewCategoriesQuery(db pgxscan.Querier) CategoriesQuery
	NewEntriesQuery(db pgxscan.Querier) EntriesQuery
	NewReportsQuery(db pgxscan.Querier) ReportsQuery
}

type queryFactory struct {
	dbRunner DbRunner
}

func NewQueryFactory(connectionPool *pgxpool.Pool) QueryFactory {
	return &queryFactory{
		dbRunner: newDbRunner(connectionPool),
	}
}

func (f *queryFactory) GetDbRunner() DbRunner {
	return f.dbRunner
}

func (f *queryFactory) NewCategoriesQuery(db pgxscan.Querier) CategoriesQuery {
	return newCategoriesQuery(db)
}

func (f *queryFactory) NewEntriesQuery(db pgxscan.Querier) EntriesQuery {
	return newEntriesQuery(db)
}

func (f *queryFactory) NewReportsQuery(db pgxscan.Querier) ReportsQuery {
	return newReportsQuery(db)
}
