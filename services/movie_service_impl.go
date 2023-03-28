package services

import (
	"context"
	"gorm.io/gorm"
	"log"
	"moviesapp/models"
	pb "moviesapp/v1"
	"time"
)

type MovieServiceImpl struct {
	movieDB *gorm.DB
	ctx     context.Context
}

func NewMovieService(movieDB *gorm.DB, ctx context.Context) MovieService {
	return &MovieServiceImpl{movieDB, ctx}
}

// implementing Create method to the above defined struct
func (m *MovieServiceImpl) CreateMovie(movie *models.Movie) (*models.Movie, error) {

	movie.CreatedAt = time.Now()
	movie.UpdatedAt = movie.CreatedAt

	err := m.movieDB.Create(movie).Error

	if err != nil {
		log.Printf("Unable to create new record; ", err)
		return nil, err
	}

	return movie, nil
}

// adding Update function to the above defined struct
func (m *MovieServiceImpl) UpdateMovie(id string, update_movie *models.Movie) (*models.Movie, error) {

	update_movie.UpdatedAt = time.Now()
	movie := &models.Movie{Isbn: id}
	m.movieDB.Model(movie).Updates(update_movie)

	return update_movie, nil
}

func (m *MovieServiceImpl) GetMovieById(id string) (*models.Movie, error) {

	//id = "7658916879"

	var movie = &models.Movie{Isbn: id}

	if err := m.movieDB.First(movie).Error; err != nil {
		log.Printf("Unable to fetch record: ", err)
		return nil, err
	}

	return movie, nil

}

func (m *MovieServiceImpl) GetMovies() ([]*models.Movie, error) {

	var all_movies []*models.Movie

	if err := m.movieDB.Find(&all_movies).Error; err != nil {
		log.Printf("Unable to fetch any record;", err)
		return nil, err
	}

	if len(all_movies) == 0 {
		return []*models.Movie{}, nil
	}

	return all_movies, nil
}

func (m *MovieServiceImpl) DeleteMovie(id string) (*pb.Status, error) {

	status := &pb.Status{
		Success: true,
	}
	success := status.GetSuccess()

	var movie *models.Movie

	if err := m.movieDB.Where(&models.Movie{Isbn: id}).First(&movie).Error; err != nil {
		success = false
		log.Printf("Unable to fetch record;", err)
	}

	m.movieDB.Delete(&movie)

	status = &pb.Status{
		Success: success,
	}
	return status, nil
}
