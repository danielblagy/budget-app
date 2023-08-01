include .env

.PHONY: docker-up
docker-up:
	docker compose up -d

.PHONY: docker-down
docker-down:
	docker compose down

.PHONY: migrate-generate
migrate-generate:
	migrate create -ext sql -dir migrations -seq $(name)

.PHONY: migrate-up
migrate-up:
	migrate -database ${DATABASE_URL} -path migrations up

.PHONY: migrate-down
migrate-down:
	migrate -database ${DATABASE_URL} -path migrations down

.PHONY: run
run:
	go run cmd/budget-app/main.go

.PHONY: test
test:
	go test -cover ./...

.PHONY: build
build:
	go build -o build/ cmd/budget-app/main.go

.PHONY: lint
lint:
	staticcheck ./...

.PHONY: setup-e2e-env
setup-e2e-env:
	make docker-down docker-up
	sleep 3
	make migrate-up

.PHONY: run-e2e-tests
run-e2e-tests:
	go test ./e2e -tags e2e