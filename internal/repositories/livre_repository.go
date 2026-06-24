package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/patrickbrouhard/mediatekdocuments-api-go/internal/models"
)

// LivreRepository gère l'accès aux données des livres.
type LivreRepository struct {
	db *sql.DB
}

// NewLivreRepository crée un repository pour les livres.
func NewLivreRepository(db *sql.DB) *LivreRepository {
	return &LivreRepository{
		db: db,
	}
}

// Lister récupère une page de livres et le nombre total de livres.
func (r *LivreRepository) Lister(
	ctx context.Context,
	parametres models.ParametresListeLivres,
) ([]models.Livre, error) {
	colonneTri, err := colonneTriLivres(parametres.Tri)
	if err != nil {
		return nil, err
	}

	ordreTri, err := ordreTriLivres(parametres.Ordre)
	if err != nil {
		return nil, err
	}

	requeteLister := fmt.Sprintf(`
		SELECT
			l.id,
			d.titre,
			l.ISBN,
			l.auteur,
			l.collection,
			COALESCE(d.image, ''),
			g.id,
			g.libelle,
			p.id,
			p.libelle,
			r.id,
			r.libelle
		FROM livre l
		JOIN document d ON d.id = l.id
		JOIN genre g ON g.id = d.idGenre
		JOIN `+"`public`"+` p ON p.id = d.idPublic
		JOIN rayon r ON r.id = d.idRayon
		ORDER BY %s %s
	`, colonneTri, ordreTri)

	lignes, err := r.db.QueryContext(ctx, requeteLister)
	if err != nil {
		return nil, fmt.Errorf("lister les livres : %w", err)
	}
	defer lignes.Close()

	// Parcourir les lignes et construire la liste des livres
	livres := make([]models.Livre, 0) // Initialiser une slice vide pour stocker les livres
	for lignes.Next() {
		var livre models.Livre

		if err := lignes.Scan(
			&livre.ID,
			&livre.Titre,
			&livre.ISBN,
			&livre.Auteur,
			&livre.Collection,
			&livre.Image,
			&livre.Genre.ID,
			&livre.Genre.Libelle,
			&livre.Public.ID,
			&livre.Public.Libelle,
			&livre.Rayon.ID,
			&livre.Rayon.Libelle,
		); err != nil {
			return nil, fmt.Errorf("lire un livre : %w", err)
		}

		livres = append(livres, livre)
	}

	if err := lignes.Err(); err != nil {
		return nil, fmt.Errorf("parcourir les livres : %w", err)
	}

	return livres, nil
}

// colonneTriLivres retourne la colonne SQL correspondante au champ de tri fourni si elle est autorisée.
func colonneTriLivres(tri string) (string, error) {
	colonnesAutorisees := map[string]string{
		"titre":      "d.titre",
		"auteur":     "l.auteur",
		"isbn":       "l.ISBN",
		"collection": "l.collection",
	}

	colonne, existe := colonnesAutorisees[tri]
	if !existe {
		return "", fmt.Errorf("champ de tri non autorisé : %s", tri)
	}

	return colonne, nil
}

// ordreTriLivres retourne l'ordre SQL correspondant au champ d'ordre fourni.
func ordreTriLivres(ordre string) (string, error) {
	ordresAutorises := map[string]string{
		"asc":  "ASC",
		"desc": "DESC",
	}

	ordreSQL, existe := ordresAutorises[ordre]
	if !existe {
		return "", fmt.Errorf("ordre de tri non autorisé : %s", ordre)
	}

	return ordreSQL, nil
}
