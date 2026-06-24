// internal/http/response/response.go
package response

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("failed to encode json response", "error", err)
	}
}

// Erreur représente une erreur renvoyée par l'API.
type Erreur struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ErreurReponse représente le format JSON d'une réponse d'erreur.
type ErreurReponse struct {
	Erreur Erreur `json:"erreur"`
}

// ErreurJSON écrit une réponse d'erreur JSON.
func ErreurJSON(
	w http.ResponseWriter,
	statut int,
	code string,
	message string,
) {
	JSON(w, statut, ErreurReponse{
		Erreur: Erreur{
			Code:    code,
			Message: message,
		},
	})
}
