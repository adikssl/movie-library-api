package dto

type CreateMovieDTO struct {
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Director    string  `json:"director"`
	Rating      float64 `json:"rating"`
	Description string  `json:"description"`
}

type UpdateMovieDTO struct {
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Director    string  `json:"director"`
	Rating      float64 `json:"rating"`
	Description string  `json:"description"`
}
