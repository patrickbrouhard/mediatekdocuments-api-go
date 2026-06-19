// internal/http/handlers/health_handler.go
package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/patrickbrouhard/mediatekdocuments-api-go/internal/http/response"
)

type HealthHandler struct {
	db *sql.DB
}

func NewHealthHandler(db *sql.DB) *HealthHandler {
	return &HealthHandler{
		db: db,
	}
}

// Health écrit une réponse JSON indiquant que le service est opérationnel.
// Paramètres :
//   - w : le writer HTTP utilisé pour envoyer la réponse au client
//   - r : la requête HTTP reçue
//
// Retour :
//   - Aucun retour direct ; la fonction écrit dans w un JSON {"status": "ok"} avec un statut 200.
func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}

// Ready vérifie la disponibilité de la base de données et retourne l'état de préparation du service.
func (h *HealthHandler) Ready(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if err := h.db.PingContext(ctx); err != nil {
		response.JSON(w, http.StatusServiceUnavailable, map[string]string{
			"status": "not_ready",
		})
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{
		"status": "ready",
	})
}
