package test

import log "github.com/inconshreveable/log15"

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=Logger --case=underscore

type Logger interface {
	log.Logger
}
