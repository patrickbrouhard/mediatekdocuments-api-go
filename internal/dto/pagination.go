package dto

const (
	PageParDefaut       = 1
	TaillePageParDefaut = 50
	TaillePageMaximum   = 100
)

type Pagination struct {
	Page           int `json:"page"`
	TaillePage     int `json:"taille_page"`
	NombreElements int `json:"nombre_elements"`
	NombrePages    int `json:"nombre_pages"`
}
