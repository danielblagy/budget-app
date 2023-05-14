# budget-app

* [Quickstart](#quickstart)
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

## Dependencies

### build
docker, docker-compose

### sql migrations
[go-migrate](https://github.com/golang-migrate/migrate) cli tool

### golang packages

## Make Commands