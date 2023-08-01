package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=DbRunner --case=underscore

type DbRunner interface {
	Begin(ctx context.Context) (pgx.Tx, error)
}

type dbRunner struct {
	connectionPool *pgxpool.Pool
}

func newDbRunner(connectionPool *pgxpool.Pool) DbRunner {
	return &dbRunner{
		connectionPool: connectionPool,
	}
}

func (r *dbRunner) Begin(ctx context.Context) (pgx.Tx, error) {
	return r.connectionPool.Begin(ctx)
}
