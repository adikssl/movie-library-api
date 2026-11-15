package dto

type MovieDTO struct {
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Director    string  `json:"director"`
	Rating      float64 `json:"rating"`
	Description string  `json:"description"`
}
