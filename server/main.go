package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/moviesapp/config"
	gapi "github.com/moviesapp/grpc-gateway"
	"github.com/moviesapp/services"
	pb "github.com/moviesapp/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

var (
	ctx          context.Context
	movieDB      *gorm.DB
	movieService services.MovieService
	wg           sync.WaitGroup
)

func startGrpcServer(movieDB *gorm.DB, movieService services.MovieService) {

	config, err := config.LoadConfig(".")

	if err != nil {
		log.Printf("cannot load config:", err)
	}

	lis, err := net.Listen("tcp", config.SERVERPORT)
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
	log.Printf("Server listening on %s", config.SERVERPORT)
	wg.Done()

}

func startRestServer() {

	config, err := config.LoadConfig(".")

	if err != nil {
		log.Printf("cannot load config:", err)
	}

	address := config.HOST + config.SERVERPORT
	port := config.GATEWAYPORT

	mux := runtime.NewServeMux()
	err = pb.RegisterMovieServiceHandlerFromEndpoint(ctx, mux, address, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Serving gateway on connection ")
	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on port ", port)
	wg.Done()
}

func main() {

	ctx = context.TODO()
	//Config file
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Printf("cannot load config:", err)
	}

	dsn := config.DBUSERNAME + ":" + config.DBPASSWORD + "@tcp" + "(" + config.DBHOST + ":" + config.DBPORT + ")/" + config.DBNAME + "?" + "parseTime=true&loc=Local"
	movieDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//movieDB.AutoMigrate(&models.Movie{})

	if err != nil {
		log.Printf("Couldn't connect to database", err, err)
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

}
