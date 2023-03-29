package services

import (
	"github.com/moviesapp/models"
	pb "github.com/moviesapp/v1"
)

// MovieService interface contains all the method signatures for our CRUD operations
type MovieService interface {
	CreateMovie(*models.Movie) (*models.Movie, error)
	UpdateMovie(string, *models.Movie) (*models.Movie, error)
	GetMovieById(string) (*models.Movie, error)
	GetMovies() ([]*models.Movie, error)
	DeleteMovie(string) (*pb.Status, error)
}
