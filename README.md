# MediaTekDocuments API Go

Réécriture en Go de l’API REST MediaTekDocuments.

## Prérequis

- Go 1.26+
- Docker et Docker Compose
- Make

## Lancer le projet avec Docker

```bash
make up
```

L’API est disponible sur :

```text
http://localhost:8080
```

Endpoints disponibles :

```text
GET /api/v1/health
GET /api/v1/ready
```

Exemples :

```bash
curl http://localhost:8080/api/v1/health
curl http://localhost:8080/api/v1/ready
```

## Commandes utiles

```bash
make test
make tidy
make up
make down
make logs
make db-shell
make reset-dev
```

## Architecture actuelle

```text
cmd/api
→ point d’entrée de l’application

internal/config
→ lecture de la configuration par variables d’environnement

internal/database
→ connexion MySQL

internal/http
→ routes, handlers et réponses JSON
```

## Variables d’environnement

Les variables disponibles sont documentées dans :

```text
.env.example
```

## Réinitialiser la base de développement

```bash
make reset-dev
```

## Références

- [Chi Documentation](https://go-chi.io/#/README)
- [Chi examples](https://github.com/go-chi/chi/tree/master/_examples)

- [How I write HTTP services in Go after 13 years](https://grafana.com/blog/how-i-write-http-services-in-go-after-13-years/)

- [Go Documentation](https://golang.org/doc/)
- [Docker Documentation](https://docs.docker.com/)
- [Make Documentation](https://www.gnu.org/software/make/manual/make.html)
