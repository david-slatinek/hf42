package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"main/model"
	"time"
)

type BookCollection struct {
	Collection *mongo.Collection
}

func (receiver BookCollection) Insert(book model.Book) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := receiver.Collection.InsertOne(ctx, book)
	return id.InsertedID.(primitive.ObjectID).Hex(), err
}
