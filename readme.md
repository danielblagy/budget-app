# budget-app

* [Introduction](#introduction)
* [Quickstart](#quickstart)
* [Documentation](#documentation)
* [Dependencies](#dependencies)
* [Make Commands](#make-commands)
* [Web Client App](#web-client-app)

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

Access full documentation on REST API is [here](documentation.md).

## Dependencies

### build
docker, docker-compose

### sql migrations
[go-migrate](https://github.com/golang-migrate/migrate) cli tool

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

## Web Client App

A web client application is currently being developed that uses this REST API to serve UI.

Check out the repository [here](https://github.com/danielblagy/budget-app-web-client).