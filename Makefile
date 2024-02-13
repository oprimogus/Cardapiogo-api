include .env
export


.PHONY: fmt

lint:
	@gofmt -s -w .

install:
	go mod tidy

docker-build:
	docker build -f build/Dockerfile -t ${APP_REPO}/${APP_NAME}:latest .

docker-up:
	docker compose -f deployments/docker-compose.yaml up -d

docker-down:
	docker compose -f deployments/docker-compose.yaml down

mock-database:
	go run scripts/populate_local_db.go

sqlc:
	sqlc generate -f internal/infra/database/sqlc/sqlc.yaml

docs:
	swag init -g cmd/main.go

test:
	go test ./... -v -cover

run:
	make docs
	go run cmd/main.go

migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir ${MIGRATION_SOURCE_URL} -seq $$name

migration-up: 
	migrate -path ${MIGRATION_SOURCE_URL} -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

migration-down: 
	migrate -path ${MIGRATION_SOURCE_URL} -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down 1

migration-fix: 
	@read -p "Enter migration version: " version; \
	migrate -path ${MIGRATION_SOURCE_URL} -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" force $$version
