package models

type Livre struct {
	ID         string
	Titre      string
	ISBN       string
	Auteur     string
	Collection string
	Image      string

	IDGenre      string
	LibelleGenre string

	IDPublic      string
	LibellePublic string

	IDRayon      string
	LibelleRayon string
}
