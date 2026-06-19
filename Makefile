.PHONY: help run test tidy build up down logs db-shell db-reset

help:
	@echo "Commandes disponibles :"
	@echo "  make run       Lance l'API localement"
	@echo "  make test      Compile et lance les tests Go"
	@echo "  make tidy      Nettoie les dépendances Go"
	@echo "  make build     Construit l'image Docker de l'API"
	@echo "  make up        Lance l'API et MySQL avec Docker Compose"
	@echo "  make down      Arrête les services Docker Compose"
	@echo "  make logs      Affiche les logs Docker Compose"
	@echo "  make db-shell  Ouvre un shell MySQL"
	@echo "  make db-reset  Réinitialise la base MySQL Docker"

run:
	go run ./cmd/api

test:
	go test ./...

tidy:
	go mod tidy

build:
	docker compose build api

up:
	docker compose up -d

down:
	docker compose down

logs:
	docker compose logs -f

db-shell:
	MYSQL_PWD=mediatek docker compose exec mysql mysql --user=mediatek mediatek86

db-reset:
	docker compose down -v
	docker compose up -d
