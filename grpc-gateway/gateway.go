package grpc_gateway

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"log"
	"moviesapp/models"
	"moviesapp/services"
	pb "moviesapp/v1"
)

type MovieServer struct {
	pb.UnimplementedMovieServiceServer
	movieDB      *gorm.DB
	movieService services.MovieService
}

func NewGRpcMovieServer(movieDB *gorm.DB, movieService services.MovieService) (*MovieServer, error) {

	movieServer := &MovieServer{
		movieDB:      movieDB,
		movieService: movieService,
	}

	return movieServer, nil
}

// GetMovies Request
func (movieServer *MovieServer) GetMovies(empty *pb.Empty, stream pb.MovieService_GetMoviesServer) error {

	allMovies, err := movieServer.movieService.GetMovies()

	if err != nil {
		log.Printf("Unable to fetch entries", err)
	}

	for _, movie := range allMovies {

		stream.Send(&pb.Movie{
			Isbn:      movie.Isbn,
			Title:     movie.Title,
			Director:  movie.Director,
			Genre:     movie.Genre,
			Rating:    movie.Rating,
			CreatedAt: timestamppb.New(movie.CreatedAt),
			UpdatedAt: timestamppb.New(movie.UpdatedAt),
		})
	}

	return nil
}

// CreateMovieRequest
func (movieServer *MovieServer) CreateMovie(ctx context.Context, request *pb.CreateMovieRequest) (*pb.Movie, error) {

	isbn := request.GetIsbn()
	if len(isbn) != 10 {
		return nil, status.Error(401, "Invalid ISBN. Needs to be 10 digits:")
	}

	movie := &models.Movie{
		Isbn:     request.GetIsbn(),
		Title:    request.GetTitle(),
		Director: request.GetDirector(),
		Genre:    request.GetGenre(),
		Rating:   request.GetRating(),
	}
	fmt.Printf("Created Movie %v", movie)

	movie, err := movieServer.movieService.CreateMovie(movie)
	if err != nil {
		log.Printf("Create Operation failed", err)
		return nil, err
	}

	var newMovie = &pb.Movie{
		Isbn:      movie.Isbn,
		Title:     movie.Title,
		Director:  movie.Director,
		Genre:     movie.Genre,
		Rating:    movie.Rating,
		CreatedAt: timestamppb.New(movie.CreatedAt),
		UpdatedAt: timestamppb.New(movie.UpdatedAt),
	}

	return newMovie, nil
}

// GetMovieRequest
func (movieServer *MovieServer) GetMovieById(ctx context.Context, request *pb.GetMovieRequest) (*pb.Movie, error) {

	movieId := request.GetIsbn()
	if len(movieId) != 10 {
		return nil, status.Error(401, "Invalid ISBN. Needs to be 10 digits:")
	}
	//log.Printf("GET Request: %v", request)
	newMovie, err := movieServer.movieService.GetMovieById(movieId)

	if err != nil {
		log.Printf("Unable to fetch record", err)
		return &pb.Movie{}, err
	}

	res := &pb.Movie{
		Isbn:      newMovie.Isbn,
		Title:     newMovie.Title,
		Director:  newMovie.Director,
		Genre:     newMovie.Genre,
		Rating:    newMovie.Rating,
		CreatedAt: timestamppb.New(newMovie.CreatedAt),
		UpdatedAt: timestamppb.New(newMovie.UpdatedAt),
	}

	return res, nil
}

// UpdateMovie Request
func (movieServer *MovieServer) UpdateMovie(ctx context.Context, request *pb.UpdateMovieRequest) (*pb.Movie, error) {

	movieId := request.GetIsbn()
	if len(movieId) != 10 {
		return nil, status.Error(401, "Invalid ISBN. Needs to be 10 digits:")
	}

	movie := &models.Movie{
		Isbn:     request.GetIsbn(),
		Title:    request.GetTitle(),
		Director: request.GetDirector(),
		Genre:    request.GetGenre(),
		Rating:   request.GetRating(),
	}

	newMovie, err := movieServer.movieService.UpdateMovie(movieId, movie)

	if err != nil {
		log.Printf("Unable to update entry", err)
	}

	res := &pb.Movie{
		Isbn:      newMovie.Isbn,
		Title:     newMovie.Title,
		Director:  newMovie.Director,
		Genre:     newMovie.Genre,
		Rating:    newMovie.Rating,
		CreatedAt: timestamppb.New(newMovie.CreatedAt),
		UpdatedAt: timestamppb.New(newMovie.UpdatedAt),
	}

	return res, nil
}

// DeleteMovie Request
func (movieServer *MovieServer) DeleteMovie(ctx context.Context, request *pb.DeleteMovieRequest) (*pb.Status, error) {

	movieId := request.GetIsbn()
	if len(movieId) != 10 {
		return nil, status.Error(401, "Invalid ISBN. Needs to be 10 digits:")
	}

	status, err := movieServer.movieService.DeleteMovie(movieId)

	if err != nil {
		log.Printf("Unable to delete entry", err)
	}

	return status, nil
}
