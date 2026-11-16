package main

import (
	"log"
	"movie-library-api/internal/handler"
	"movie-library-api/internal/repository"
	"movie-library-api/internal/service"
	"net/http"
)

func main() {
	movieRepository := repository.NewMovieRepository()
	movieService := service.NewMovieService(movieRepository)
	movieHandler := handler.NewMovieHandler(movieService)
	movieRouter := handler.NewMovieRouter(movieHandler)

	if err := http.ListenAndServe(":8080", movieRouter); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
