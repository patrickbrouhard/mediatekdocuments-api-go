package http

import (
	"database/sql"
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/patrickbrouhard/mediatekdocuments-api-go/internal/http/handlers"
	"github.com/patrickbrouhard/mediatekdocuments-api-go/internal/http/middleware"
	"github.com/patrickbrouhard/mediatekdocuments-api-go/internal/repositories"
	"github.com/patrickbrouhard/mediatekdocuments-api-go/internal/services"
)

func NewRouter(db *sql.DB, logger *slog.Logger) *chi.Mux { // *chi.Mux = pointeur vers le routeur principal de Chi.
	r := chi.NewRouter()

	r.Use(middleware.RequestLogger(logger))

	r.Route("/api/v1", func(r chi.Router) { // préfixe /api/v1 automatique (sous-routeur)

		healthHandler := handlers.NewHealthHandler(db)

		livreRepository := repositories.NewLivreRepository(db)
		livreService := services.NewLivreService(livreRepository)
		livreHandler := handlers.NewLivreHandler(livreService)

		r.Get("/health", healthHandler.Health)
		r.Get("/ready", healthHandler.Ready)

		r.Get("/livres", livreHandler.Lister)
	})

	return r
}
