package server

import (
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"main/db"
	pb "main/schema"
)

type Server struct {
	pb.BookServiceServer
	Collection db.BookCollection
}

func (server Server) ValidateBooks(request *pb.ValidateBooksRequest, stream pb.BookService_ValidateBooksServer) error {
	if len(request.BooksISBN) == 0 {
		err := stream.Send(&pb.ValidateBooksResponse{
			Valid: false,
			Code:  int32(codes.InvalidArgument),
			Error: "no books to validate",
		})

		if err != nil {
			log.Printf("error while sending response: %v\n", err)
		}

		return status.Error(codes.InvalidArgument, "no books to validate")
	}

	log.Println("books to validate: ", request.BooksISBN)

	for _, isbn := range request.BooksISBN {
		_, err := server.Collection.GetBookByISBN(isbn)

		if errors.Is(err, mongo.ErrNoDocuments) {
			err := stream.Send(&pb.ValidateBooksResponse{
				Valid: false,
				Code:  int32(codes.NotFound),
				Error: "book with isbn=" + isbn + " not found",
			})

			if err != nil {
				log.Printf("error while sending response: %v\n", err)
			}
			continue
		}

		if err != nil {
			err := stream.Send(&pb.ValidateBooksResponse{
				Valid: false,
				Code:  int32(codes.Internal),
				Error: err.Error(),
			})
			if err != nil {
				log.Printf("error while sending response: %v\n", err)
			}
			continue
		}

		err = stream.Send(&pb.ValidateBooksResponse{
			Valid: true,
			Code:  int32(codes.OK),
			Error: "",
		})
		if err != nil {
			log.Printf("error while sending response: %v\n", err)
		}
	}

	return nil
}
