package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"main/controller"
	"main/db"
	_ "main/docs"
	"main/env"
	pb "main/schema"
	"main/server"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//	@title			Book management API
//	@version		1.0
//	@description	Book management API for hf42 project
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	David Slatinek
//	@contact.url	https://github.com/david-slatinek

//	@accept		json
//	@produce	json
//	@schemes	http

//	@license.name	GNU General Public License v3.0
//	@license.url	https://www.gnu.org/licenses/gpl-3.0.html

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

// @host	localhost:8080
func main() {
	err := env.Load("env/.env")
	if err != nil {
		log.Fatalf("failed to load env: %v", err)
	}

	uri := os.Getenv("MONGO_URL")
	if uri == "" {
		log.Fatal("MONGO_URL is not set")
	}

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("error with Connect: %v", err)
	}
	defer func(client *mongo.Client) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := client.Disconnect(ctx); err != nil {
			log.Printf("error with Disconnect: %v", err)
		}
	}(client)

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("error with Ping: %v", err)
	}

	bookController := controller.BookController{
		Collection: db.BookCollection{
			Collection: client.Database("book-service").Collection("books"),
		},
	}

	router := gin.Default()

	router.POST("/book", bookController.CreateBook)
	router.GET("/book/:isbn", bookController.GetBookByISBN)
	router.PUT("/book", bookController.UpdateBook)
	router.DELETE("/book/:isbn", bookController.DeleteBookByISBN)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	srv := &http.Server{
		Addr:         ":8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("ListenAndServe() error: %s\n", err)
		}
	}()

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000 with error: %v\n", err)
	}

	grpcServer := grpc.NewServer()

	mainGrpcServer := server.Server{
		Collection: db.BookCollection{
			Collection: client.Database("book-service").Collection("books"),
		},
	}

	pb.RegisterBookServiceServer(grpcServer, mainGrpcServer)
	reflection.Register(grpcServer)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve, error: %v", err)
		}
	}()

	<-c

	grpcServer.GracefulStop()
	log.Println("shutting down gRPC server")

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Shutdown() error: %s\n", err)
	}
	log.Println("shutting down")
}
