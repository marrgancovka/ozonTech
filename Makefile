include .env

run:
    go run cmd/main/main.go -storage=in-memory

run-postgres:
    go run cmd/main/main.go -storage=postgres

build:
    docker-compose up --build

test:
    go test ./...


.PHONY: lint
lint:
	golangci-lint run --config=.golangci.yaml

test:
	go test -race ./...

migrate-lib:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate-up: migrate-lib
	migrate -path schema -database "postgres://$(DB_USER):$(DB_PASS)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

migrate-down:
	migrate -path schema -database "postgres://$(DB_USER):$(DB_PASS)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" down

compose-build:
	docker-compose build

compose-up:
	docker-compose up -d

compose-down:
	docker-compose down