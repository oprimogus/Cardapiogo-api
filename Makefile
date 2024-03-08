include .env
export


.PHONY: fmt lint install docker-up docker-down dev-docker-up dev-docker-down mock-database sqlc docs test run

lint:
	@gofmt -s -w .

install:
	go mod tidy

docker-up:
	docker compose -f deployments/docker-compose.yaml up --build -d

docker-down:
	docker compose -f deployments/docker-compose.yaml down

dev-docker-up:
	docker compose -f deployments/docker-compose-development.yaml up -d

dev-docker-down:
	docker compose -f deployments/docker-compose-development.yaml down

mock-database:
	go run scripts/populate_local_db.go

mock:
	@read -p "Enter mock interface path: " name; \
		mockgen -source=internal/domain/$$name/repository.go -destination=internal/infra/mocks/$$name/mock_$$name.go 

sqlc:
	sqlc generate -f configs/sqlc.yaml

docs:
	swag init -g cmd/main.go

test:
	go test ./... -v -cover

run:
	make docs
	go run cmd/main.go

migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir internal/infra/database/migrations -seq $$name

migration-up: 
	migrate -path internal/infra/database/migrations -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

migration-down: 
	migrate -path internal/infra/database/migrations -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down 1

migration-fix: 
	@read -p "Enter migration version: " version; \
	migrate -path internal/infra/database/migrations -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" force $$version
