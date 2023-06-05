# budget-app

* [Quickstart](#quickstart)
* [Documentation](#documentation)
* [Dependencies](#dependencies)
* [Make Commands](#make-commands)

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

### golang packages

TODO

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