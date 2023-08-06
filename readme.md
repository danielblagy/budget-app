# budget-app

* [Introduction](#introduction)
* [Quickstart](#quickstart)
* [Documentation](#documentation)
* [Dependencies](#dependencies)
* [Make Commands](#make-commands)
* [Web Client App](#web-client-app)
* [E2E Tests](#e2e-tests)

## Introduction 

Budget-app is a personal finance app (cash organizer) with the following functionality: 
- Register in the application
- Log in to the application (if you have an account)
- View/add/edit/delete *Categories*
- View/add/edit/delete *Expenses*
- View/add/edit/delete *Incomes*
- View Reports

## Quickstart

Start Docker containers
```
make docker-up
```

Run database migrations
```
make migrate-up
```

Run REST server
```
make run
```

## Documentation

Access full documentation on REST API [here](documentation.md).

## Dependencies

### build
docker, docker-compose

### sql migrations
[go-migrate](https://github.com/golang-migrate/migrate) cli tool

### database

PostgreSQL

### fast-access persistent storage

Redis

Redis is used as a fast-access persistent storage for handling invalid JWT tokens.

## Make Commands

Start Docker containers
```
make docker-up
```

Stop and remove Docker containers
```
make docker-down
```

Create a database migration
```
make migrate-generate name=some_name_for_your_migration
```

Run database migrations up
```
make migrate-up
```

Run database migrations down
```
make migrate-down
```

Run REST server
```
make run
```

Run tests
```
make test
```

Run linter
```
make lint
```

Build project (builds to directory /build)
```
make build
```

## Web Client App

A web client application is currently being developed that uses this REST API to serve UI.

Check out the repository [here](https://github.com/danielblagy/budget-app-web-client).

## E2E Tests

E2E Tests are run on each pull request into master via Github Actions.

To run e2e tests locally:
1. `make setup-e2e-env` to set up the testing environment (this simply shuts down and removes running containers if they exist, starts docker containers, runs postgresql migrations)
2. `make run` to starts the REST server (for convenience, use a separate terminal window for this)
3. `make run-e2e-tests` to run e2e tests

E2E testing is done using go test. Unit tests and e2e tests are separated and will not run simultaneously.