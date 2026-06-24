package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/patrickbrouhard/mediatekdocuments-api-go/internal/models"
	"github.com/patrickbrouhard/mediatekdocuments-api-go/internal/repositories"
)

// ErrLivreIntrouvable indique qu'aucun livre ne correspond à l'identifiant demandé.
var ErrLivreIntrouvable = errors.New("livre introuvable")

// LivreService contient la logique applicative liée aux livres.
type LivreService struct {
	repository *repositories.LivreRepository
}

// NewLivreService crée un service pour les livres.
func NewLivreService(repository *repositories.LivreRepository) *LivreService {
	return &LivreService{
		repository: repository,
	}
}

// Lister récupère tous les livres.
func (s *LivreService) Lister(
	ctx context.Context,
	parametres models.ParametresListeLivres,
) ([]models.Livre, error) {
	livres, err := s.repository.Lister(ctx, parametres)
	if err != nil {
		return nil, err
	}

	return livres, nil
}

// RecupererParID récupère un livre à partir de son identifiant.
func (s *LivreService) RecupererParID(
	ctx context.Context,
	id string,
) (models.Livre, error) {
	livre, err := s.repository.RecupererParID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Livre{}, ErrLivreIntrouvable
		}

		return models.Livre{}, err
	}

	return livre, nil
}
