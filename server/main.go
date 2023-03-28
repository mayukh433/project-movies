package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	gapi "moviesapp/grpc-gateway"
	"moviesapp/services"
	pb "moviesapp/v1"
	"net"
	"net/http"
	"sync"
	"time"
)

const (
	port = ":8000"
)

var (
	ctx          context.Context
	movieDB      *gorm.DB
	movieService services.MovieService
	wg           sync.WaitGroup
)

const (
	DB_USERNAME = "root"
	DB_PASSWORD = "test123"
	DB_HOST     = "localhost"
	DB_PORT     = "3306"
	DB_NAME     = "grpcproj"
)

func startGrpcServer(movieDB *gorm.DB, movieService services.MovieService) {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Printf("failed to listen")
	}

	s := grpc.NewServer()

	movieServer, err := gapi.NewGRpcMovieServer(movieDB, movieService)

	if err != nil {
		log.Printf("Failed to create new server", err)
	}

	pb.RegisterMovieServiceServer(s, movieServer)

	log.Println("Serving gRPC on connection ")

	go func() {
		//defer wg.Done()
		err = s.Serve(lis)
		if err != nil {
			fmt.Printf("Failed to list server", err)
		}
	}()
	log.Printf("Server listening on localhost:8000")
	wg.Done()

}

func startRestServer() {

	mux := runtime.NewServeMux()

	err := pb.RegisterMovieServiceHandlerFromEndpoint(ctx, mux, "localhost:8000", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatal(err)
	}

	port := ":8081"

	log.Println("Serving gateway on connection ")
	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening on port 8081")
	wg.Done()
}

func main() {

	ctx = context.TODO()
	//Config file
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	movieDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//movieDB.AutoMigrate(&models.Movie{})

	if err != nil {
		log.Printf("Couldn't connect to database", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	movieService := services.NewMovieService(movieDB, ctx)

	wg.Add(2)
	// Initializing a grpc server

	startGrpcServer(movieDB, movieService)
	// Initializing a REST server
	startRestServer()
	wg.Wait()
	//time.Sleep(10 * time.Second)

}
