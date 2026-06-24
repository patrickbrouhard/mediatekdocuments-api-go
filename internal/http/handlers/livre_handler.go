package handlers

import (
	"log/slog"
	"net/http"

	"github.com/patrickbrouhard/mediatekdocuments-api-go/internal/http/response"
	"github.com/patrickbrouhard/mediatekdocuments-api-go/internal/models"
	"github.com/patrickbrouhard/mediatekdocuments-api-go/internal/services"
)

// LivreHandler gère les requêtes HTTP liées aux livres.
type LivreHandler struct {
	service *services.LivreService
}

// NewLivreHandler crée un handler pour les livres.
func NewLivreHandler(service *services.LivreService) *LivreHandler {
	return &LivreHandler{
		service: service,
	}
}

type listeLivresReponse struct {
	Donnees []models.Livre `json:"donnees"`
}

// Lister retourne la liste des livres avec les paramètres de tri et d'ordre.
func (h *LivreHandler) Lister(w http.ResponseWriter, r *http.Request) {
	parametres, err := lireParametresListeLivres(r)
	if err != nil {
		response.ErreurJSON(
			w,
			http.StatusBadRequest,
			"parametres_invalides",
			err.Error(),
		)
		return
	}

	livres, err := h.service.Lister(r.Context(), parametres)
	if err != nil {
		slog.Error(
			"échec de la récupération des livres",
			"error", err,
		)

		response.ErreurJSON(
			w,
			http.StatusInternalServerError,
			"erreur_interne",
			"Une erreur interne est survenue.",
		)
		return
	}

	response.JSON(w, http.StatusOK, listeLivresReponse{
		Donnees: livres,
	})
}

func lireParametresListeLivres(
	r *http.Request,
) (models.ParametresListeLivres, error) {
	tri := r.URL.Query().Get("tri")
	if tri == "" {
		tri = "titre"
	}

	if !triLivreAutorise(tri) {
		return models.ParametresListeLivres{},
			erreurParametre("tri", "n'est pas autorisé")
	}

	ordre := r.URL.Query().Get("ordre")
	if ordre == "" {
		ordre = "asc"
	}

	if ordre != "asc" && ordre != "desc" {
		return models.ParametresListeLivres{},
			erreurParametre("ordre", "doit être asc ou desc")
	}

	return models.ParametresListeLivres{
		Tri:   tri,
		Ordre: ordre,
	}, nil
}

func triLivreAutorise(tri string) bool {
	trisAutorises := map[string]bool{
		"titre":      true,
		"auteur":     true,
		"isbn":       true,
		"collection": true,
	}

	return trisAutorises[tri]
}

func erreurParametre(nom, message string) error {
	return &erreurValidation{
		message: "Le paramètre " + nom + " " + message + ".",
	}
}

type erreurValidation struct {
	message string
}

func (e *erreurValidation) Error() string {
	return e.message
}
