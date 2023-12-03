install:
	go mod tidy

run-database:
	docker compose -f docker-compose.yaml up -d

run:
	go run main.go

migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir ./migrations -seq $$name 