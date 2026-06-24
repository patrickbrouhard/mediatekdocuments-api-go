package services

import (
	"context"

	"github.com/patrickbrouhard/mediatekdocuments-api-go/internal/models"
	"github.com/patrickbrouhard/mediatekdocuments-api-go/internal/repositories"
)

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
