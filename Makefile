.SILENT:
-include .env

DB_URL=postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=disable

run:
	go run cmd/main.go

print:
	echo "$(DB_URL)"

test: 
	go test -v -cover ./storage/postgres

migrate-up:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" -verbose down

swag-init:
	swag init -g api/server.go -o api/docs

compose-up:
	docker compose --env-file .env.docker up

compose-down:
	docker compose down