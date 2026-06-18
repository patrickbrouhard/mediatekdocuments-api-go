APP_NAME=mediatekdocuments-api-go

.PHONY: help run test tidy docker-up docker-down docker-logs db-shell db-reset

help:
	@echo "Available commands:"
	@echo "  make run          Run the API locally"
	@echo "  make test         Run Go tests"
	@echo "  make tidy         Tidy Go modules"
	@echo "  make docker-up    Start Docker services"
	@echo "  make docker-down  Stop Docker services"
	@echo "  make docker-logs  Show Docker logs"
	@echo "  make db-shell     Open MySQL shell"
	@echo "  make db-reset     Reset Docker MySQL volume"

run:
	go run ./cmd/api

test:
	go test ./...

tidy:
	go mod tidy

docker-up:
	docker compose up -d

docker-down:
	docker compose down

docker-logs:
	docker compose logs -f

db-shell:
	docker compose exec mysql mysql --user=mediatek --password=mediatek mediatek86

db-reset:
	docker compose down -v
	docker compose up -d mysql
