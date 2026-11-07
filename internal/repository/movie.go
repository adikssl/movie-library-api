package repository

import (
	"movie-library-api/internal/model"
	"sync"
)

type MovieRepository struct {
	mtx    sync.RWMutex
	movies map[int]model.Movie
	nextID int
}

func NewMovieRepository() *MovieRepository {
	return &MovieRepository{
		movies: make(map[int]model.Movie),
		nextID: 1,
	}
}

func (m *MovieRepository) CreateMovie(movie model.Movie) model.Movie {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	movie.ID = m.nextID
	m.nextID += 1
	m.movies[movie.ID] = movie
	return movie
}

func (m *MovieRepository) GetMovieByID(id int) (model.Movie, bool) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	task, ok := m.movies[id]
	if !ok {
		return model.Movie{}, false
	}
	return task, true
}

func (m *MovieRepository) GetAllMovies() []model.Movie {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	allMovies := []model.Movie{}
	for _, movie := range m.movies {
		allMovies = append(allMovies, movie)
	}
	return allMovies
}

func (m *MovieRepository) UpdateMovie(movie model.Movie) bool {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	if _, ok := m.movies[movie.ID]; !ok {
		return false
	}
	m.movies[movie.ID] = movie
	return true
}

func (m *MovieRepository) DeleteMovie(id int) bool {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	_, ok := m.movies[id]
	if !ok {
		return false
	}
	delete(m.movies, id)
	return true
}
