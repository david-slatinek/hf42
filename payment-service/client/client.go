package client

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"main/model"
	pb "main/schema"
	"os"
	"time"
)

type Client struct {
	connection *grpc.ClientConn
	client     pb.BookServiceClient
}

func NewClient() (Client, error) {
	conn, err := grpc.Dial(os.Getenv("BOOK_URL")+":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return Client{}, err
	}
	return Client{
		connection: conn,
		client:     pb.NewBookServiceClient(conn),
	}, nil
}

func (receiver Client) Close() error {
	return receiver.connection.Close()
}

func (receiver Client) ValidateBooks(order model.Order) error {
	booksISBN := make([]string, len(order.Books))
	for i, book := range order.Books {
		booksISBN[i] = book.ISBN
	}

	request := &pb.ValidateBooksRequest{
		BooksISBN: booksISBN,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := receiver.client.ValidateBooks(ctx, request)
	if err != nil {
		return err
	}

	valid := make([]bool, 0, len(order.Books))
	errs := make([]string, 0, len(order.Books))

	for {
		response, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if response.Valid == false || response.Code != 0 || response.Error != "" {
			return errors.New(response.Error)
		}

		valid = append(valid, response.Valid)
		errs = append(errs, response.Error)
	}

	for k, v := range valid {
		if v == false {
			return errors.New(errs[k])
		}
	}

	return nil
}
