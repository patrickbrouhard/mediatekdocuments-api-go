package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/patrickbrouhard/mediatekdocuments-api-go/internal/http/handlers"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", handlers.Health)
	})

	return r
}
