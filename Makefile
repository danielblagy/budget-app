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
	go test ./...