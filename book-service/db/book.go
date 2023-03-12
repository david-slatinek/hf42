package db

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"main/model"
	"time"
)

type BookCollection struct {
	Collection *mongo.Collection
}

func (receiver BookCollection) Insert(book model.Book) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	book2, err := receiver.GetBookByISBN(book.ISBN)
	if err == nil && book2.ISBN == book.ISBN {
		return errors.New("book with isbn=" + book.ISBN + " already exists")
	}

	_, err = receiver.Collection.InsertOne(ctx, book)
	return err
}

func (receiver BookCollection) GetBookByISBN(isbn string) (model.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var book model.Book
	if err := receiver.Collection.FindOne(ctx, bson.M{"isbn": isbn}).Decode(&book); err != nil {
		return model.Book{}, err
	}
	return book, nil
}

func (receiver BookCollection) UpdateBook(book model.Book) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := receiver.Collection.ReplaceOne(ctx, bson.M{"isbn": book.ISBN}, book)
	return int(res.ModifiedCount), err
}

func (receiver BookCollection) DeleteBookByISBN(isbn string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := receiver.Collection.DeleteOne(ctx, bson.M{"isbn": isbn})
	return int(res.DeletedCount), err
}
