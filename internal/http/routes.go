package http

import (
	"database/sql"
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/patrickbrouhard/mediatekdocuments-api-go/internal/http/handlers"
	"github.com/patrickbrouhard/mediatekdocuments-api-go/internal/http/middleware"
)

func NewRouter(db *sql.DB, logger *slog.Logger) *chi.Mux { // *chi.Mux = pointeur vers le routeur principal de Chi.
	r := chi.NewRouter()

	r.Use(middleware.RequestLogger(logger))

	r.Route("/api/v1", func(r chi.Router) { // préfixe /api/v1 automatique (sous-routeur)

		healthHandler := handlers.NewHealthHandler(db)

		r.Get("/health", healthHandler.Health)
		r.Get("/ready", healthHandler.Ready)
	})

	return r
}
