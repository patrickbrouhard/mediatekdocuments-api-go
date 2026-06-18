package handlers

import (
	"encoding/json"
	"net/http"
)

// Health écrit une réponse JSON indiquant que le service est opérationnel.
// Paramètres :
//   - w : le writer HTTP utilisé pour envoyer la réponse au client
//   - r : la requête HTTP reçue
//
// Retour :
//   - Aucun retour direct ; la fonction écrit dans w un JSON {"status": "ok"} avec un statut 200.
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

// Ready doit ensuite être changé quand on aura connecté la base de données et que l'on pourra vérifier si elle est prête. Pour l'instant, on renvoie toujours "ready".
func Ready(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "ready",
	})
}
