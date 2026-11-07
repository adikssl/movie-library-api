package model

type Movie struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Director    string  `json:"director"`
	Rating      float64 `json:"rating"`
	Description string  `json:"description"`
}
