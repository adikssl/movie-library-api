package service

import "errors"

var ErrMovieNotFound = errors.New("Movie not found")
var ErrInvalidTitle = errors.New("Invalid title")
var ErrInvalidYear = errors.New("Invalid year")
var ErrInvalidRating = errors.New("Invalid rating")
var ErrInvalidDirector = errors.New("Invalid director")
