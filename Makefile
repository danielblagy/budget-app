include .env


migrate-up:
	migrate -database ${DATABASE_URL} -path migrations up

migrate-down:
	migrate -database ${DATABASE_URL} -path migrations down

run:
	go run cmd/budget-app/main.go