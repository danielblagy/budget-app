include .env

docker-up:
	docker compose up -d

docker-down:
	docker compose down

migrate-generate:
	migrate create -ext sql -dir migrations -seq $(name)

migrate-up:
	migrate -database ${DATABASE_URL} -path migrations up

migrate-down:
	migrate -database ${DATABASE_URL} -path migrations down

run:
	go run cmd/budget-app/main.go

test:
	go test ./...