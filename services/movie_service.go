package services

import (
	"moviesapp/models"
	pb "moviesapp/v1"
)

type MovieService interface {
	CreateMovie(*models.Movie) (*models.Movie, error)
	UpdateMovie(string, *models.Movie) (*models.Movie, error)
	GetMovieById(string) (*models.Movie, error)
	GetMovies() ([]*models.Movie, error)
	DeleteMovie(string) (*pb.Status, error)
}

type DummyService interface {
}
