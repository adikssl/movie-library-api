package handler

import "github.com/gorilla/mux"

func NewMovieRouter(movieHandler *MovieHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/movies", movieHandler.CreateMovieHandler).Methods("POST")
	router.HandleFunc("/movies", movieHandler.GetAllMoviesHandler).Methods("GET")
	router.HandleFunc("/movies/{id}", movieHandler.GetMovieByIDHandler).Methods("GET")
	router.HandleFunc("/movies/{id}", movieHandler.UpdateMovieHandler).Methods("PUT")
	router.HandleFunc("/movies/{id}", movieHandler.DeleteMovieHandler).Methods("DELETE")

	return router
}
