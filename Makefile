include .env

DOCKER_COMPOSE_FILE=docker-compose.yaml
IN_MEMORY_ENV_FILE=.env.inmemory
POSTGRES_ENV_FILE=.env.postgres

test:
	go test -race ./...

migrate-lib:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate-up: migrate-lib
	migrate -path schema -database "postgres://$(DB_USER):$(DB_PASS)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

migrate-down:
	migrate -path schema -database "postgres://$(DB_USER):$(DB_PASS)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" down

start:
	docker compose build
	docker compose up -d

start:
	docker compose build
	docker compose up -d

stop:
	docker-compose down
