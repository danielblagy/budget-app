package test

import (
	log "github.com/inconshreveable/log15"
	"github.com/jackc/pgx/v5"
)

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=Logger --case=underscore
//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=Tx --case=underscore

type Logger interface {
	log.Logger
}

type Tx interface {
	pgx.Tx
}
