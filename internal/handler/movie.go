package handler

import (
	"encoding/json"
	"errors"
	"movie-library-api/internal/handler/dto"
	"movie-library-api/internal/model"
	"movie-library-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MovieService interface {
	CreateMovie(movie model.Movie) (model.Movie, error)
	GetMovieByID(id int) (model.Movie, error)
	GetAllMovies() []model.Movie
	UpdateMovie(movie model.Movie) error
	DeleteMovie(id int) error
}

type MovieHandler struct {
	service MovieService
}

func NewMovieHandler(service MovieService) *MovieHandler {
	return &MovieHandler{
		service: service,
	}
}

func getStatusCode(err error) int {
	switch {
	case errors.Is(err, service.ErrInvalidTitle),
		errors.Is(err, service.ErrInvalidYear),
		errors.Is(err, service.ErrInvalidRating),
		errors.Is(err, service.ErrInvalidDirector):
		return http.StatusBadRequest

	case errors.Is(err, service.ErrMovieNotFound):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func (m *MovieHandler) writeError(
	w http.ResponseWriter,
	err error,
	status int,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(dto.NewError(err))
}

func (m *MovieHandler) CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	var createMovieDTO dto.MovieDTO
	if err := json.NewDecoder(r.Body).Decode(&createMovieDTO); err != nil {
		m.writeError(w, err, http.StatusBadRequest)
		return
	}
	movie := model.Movie{
		Title:       createMovieDTO.Title,
		Description: createMovieDTO.Description,
		Director:    createMovieDTO.Director,
		Year:        createMovieDTO.Year,
		Rating:      createMovieDTO.Rating,
	}
	movie, err := m.service.CreateMovie(movie)
	if err != nil {
		m.writeError(w, err, getStatusCode(err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

func (m *MovieHandler) GetMovieByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		m.writeError(w, err, http.StatusBadRequest)
		return
	}
	movie, err := m.service.GetMovieByID(id)
	if err != nil {
		m.writeError(w, err, getStatusCode(err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}

func (m *MovieHandler) GetAllMoviesHandler(w http.ResponseWriter, r *http.Request) {
	allMovies := m.service.GetAllMovies()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allMovies)
}

func (m *MovieHandler) UpdateMovieHandler(w http.ResponseWriter, r *http.Request) {
	var updateMovieDTO dto.MovieDTO
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		m.writeError(w, err, http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&updateMovieDTO); err != nil {
		m.writeError(w, err, http.StatusBadRequest)
		return
	}
	movie := model.Movie{
		ID:          id,
		Title:       updateMovieDTO.Title,
		Description: updateMovieDTO.Description,
		Director:    updateMovieDTO.Director,
		Year:        updateMovieDTO.Year,
		Rating:      updateMovieDTO.Rating,
	}
	err = m.service.UpdateMovie(movie)
	if err != nil {
		m.writeError(w, err, getStatusCode(err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}

func (m *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		m.writeError(w, err, http.StatusBadRequest)
		return
	}
	err = m.service.DeleteMovie(id)
	if err != nil {
		m.writeError(w, err, getStatusCode(err))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
