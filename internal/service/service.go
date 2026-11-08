package service

import (
	"movie-library-api/internal/model"
	"time"
)

type MovieRepository interface {
	CreateMovie(movie model.Movie) model.Movie
	GetMovieByID(id int) (model.Movie, bool)
	GetAllMovies() []model.Movie
	UpdateMovie(movie model.Movie) bool
	DeleteMovie(id int) bool
}

type MovieService struct {
	repo MovieRepository
}

func NewMovieService(repo MovieRepository) *MovieService {
	return &MovieService{
		repo: repo,
	}
}

func (m *MovieService) validateMovie(movie model.Movie) error {
	if movie.Title == "" {
		return ErrInvalidTitle
	}
	if movie.Year <= 1888 && movie.Year >= time.Now().Year() {
		return ErrInvalidYear
	}
	if movie.Rating < 0 || movie.Rating > 10 {
		return ErrInvalidRating
	}
	if movie.Director == "" {
		return ErrInvalidDirector
	}
	return nil
}

func (m *MovieService) CreateMovie(movie model.Movie) (model.Movie, error) {
	if err := m.validateMovie(movie); err != nil {
		return model.Movie{}, err
	}
	return m.repo.CreateMovie(movie), nil
}

func (m *MovieService) GetMovieByID(id int) (model.Movie, error) {
	movie, ok := m.repo.GetMovieByID(id)
	if !ok {
		return model.Movie{}, ErrMovieNotFound
	}
	return movie, nil
}

func (m *MovieService) GetAllMovies() []model.Movie {
	return m.repo.GetAllMovies()
}

func (m *MovieService) UpdateMovie(movie model.Movie) error {
	if err := m.validateMovie(movie); err != nil {
		return err
	}

	ok := m.repo.UpdateMovie(movie)
	if !ok {
		return ErrMovieNotFound
	}

	return nil
}

func (m *MovieService) DeleteMovie(id int) error {
	ok := m.repo.DeleteMovie(id)
	if !ok {
		return ErrMovieNotFound
	}
	return nil
}
