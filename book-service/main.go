package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"main/controller"
	"main/db"
	"main/env"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

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

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-c

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Shutdown() error: %s\n", err)
	}
	log.Println("shutting down")
}
