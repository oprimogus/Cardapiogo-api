include .env
export


.PHONY: fmt lint install up down stop mock-database sqlc docs test run

lint:
	@gofmt -s -w .

install:
	go mod tidy

up:
	docker compose -f deployments/docker-compose.yaml up -d

down:
	docker compose -f deployments/docker-compose.yaml down

stop:
	docker compose -f deployments/docker-compose.yaml stop

mock-db:
	go run scripts/populate_local_db.go

sqlc:
	sqlc generate -f configs/sqlc.yaml

sqlc-vet:
	sqlc vet -f configs/sqlc.yaml

docs:
	make lint
	swag init -g cmd/main.go -o api 

test:
	go test ./... -v -cover -coverprofile=cover.out

test-cover:
	go tool cover -html=cover.out

dev:
	make docs
	air 

run:
	make docs
	go run cmd/main.go

migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir internal/infrastructure/database/migrations -seq $$name

migration-up: 
	migrate -path internal/infrastructure/database/migrations -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable&search_path=public" -verbose up

migration-down: 
	migrate -path internal/infrastructure/database/migrations -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable&search_path=public" -verbose down

migration-fix: 
	@read -p "Enter migration version: " version; \
	migrate -path internal/infrastructure/database/migrations -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable&search_path=public" force $$version
