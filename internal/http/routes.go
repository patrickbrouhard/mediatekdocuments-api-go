package http

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/patrickbrouhard/mediatekdocuments-api-go/internal/http/handlers"
)

func NewRouter(db *sql.DB) *chi.Mux { // *chi.Mux = pointeur vers le routeur principal de Chi.
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) { // préfixe /api/v1 automatique (sous-routeur)

		healthHandler := handlers.NewHealthHandler(db)

		r.Get("/health", healthHandler.Health)
		r.Get("/ready", healthHandler.Ready)
	})

	return r
}
