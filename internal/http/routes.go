package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/patrickbrouhard/mediatekdocuments-api-go/internal/http/handlers"
)

func NewRouter() *chi.Mux { // *chi.Mux = pointeur vers le routeur principal de Chi.
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) { // préfixe /api/v1 automatique (sous-routeur)
		r.Get("/health", handlers.Health)
		r.Get("/ready", handlers.Ready)
	})

	return r
}
