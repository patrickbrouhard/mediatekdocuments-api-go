package dto

type ElementLibelle struct {
	ID      string `json:"id"`
	Libelle string `json:"libelle"`
}

type LivreReponse struct {
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

type ListeLivresReponse struct {
	Donnees    []LivreReponse `json:"donnees"`
	Pagination Pagination     `json:"pagination"`
}
