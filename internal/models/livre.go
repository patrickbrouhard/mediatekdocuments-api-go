package models

// ElementLibelle représente un élément avec un identifiant et un libellé.
type ElementLibelle struct {
	ID      string `json:"id"`
	Libelle string `json:"libelle"`
}

// Livre représente un livre dans la base de données tel qu'il sera renvoyé par l'API.
type Livre struct {
	ID         string `json:"id"`
	Titre      string `json:"titre"`
	ISBN       string `json:"isbn"`
	Auteur     string `json:"auteur"`
	Collection string `json:"collection"`
	Image      string `json:"image"`

	Genre  ElementLibelle `json:"genre"`
	Public ElementLibelle `json:"public"`
	Rayon  ElementLibelle `json:"rayon"`
}

// ParametresListeLivres représente les paramètres pour la liste des livres
type ParametresListeLivres struct {
	Tri   string
	Ordre string
}
