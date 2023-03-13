package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main/db"
	"main/env"
	"main/model"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var testBook = model.Book{
	ISBN:             "test-isbn",
	Title:            "test-title",
	Subtitle:         "test-subtitle",
	Author:           "test-author",
	Year:             time.Now().Year(),
	Description:      "test-description",
	Categories:       []string{"test-category"},
	OriginalTitle:    "test-original-title",
	OriginalSubtitle: "test-original-subtitle",
	OriginalYear:     time.Now().Year() - 1,
	Translator:       "test-translator",
	Size:             "test-size",
	Weight:           "test-weight",
	Pages:            420,
	Publisher:        "test-publisher",
	Language:         "test-language",
	Price:            42,
}

func getClient() *mongo.Client {
	_ = env.Load("../env/.env")

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL")).SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, _ := mongo.Connect(ctx, clientOptions)
	return client
}

func getBookController() (BookController, func()) {
	client := getClient()
	return BookController{
			Collection: db.BookCollection{
				Collection: client.Database("book-service").Collection("books"),
			},
		}, func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			_ = client.Disconnect(ctx)
		}
}

func TestBookController_CreateBook(t *testing.T) {
	bookController, tearDown := getBookController()
	defer tearDown()
	defer func(Collection db.BookCollection, isbn string) {
		_, _ = Collection.DeleteBookByISBN(isbn)
	}(bookController.Collection, testBook.ISBN)

	jsonReq, _ := json.Marshal(testBook)
	req, _ := http.NewRequest(http.MethodPost, "/book", bytes.NewBuffer(jsonReq))

	router := gin.Default()
	router.POST("/book", bookController.CreateBook)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Logf("Expected return code was %d, got %d", http.StatusCreated, w.Code)

		var errDesc map[string]string
		if err := json.Unmarshal(w.Body.Bytes(), &errDesc); err != nil {
			t.Logf("Error with unmarshal: %s", err)
			t.FailNow()
		}

		if val, ok := errDesc["error"]; ok {
			t.Logf("Error description: %s", val)
		} else {
			t.Logf("Error description not found, response: %s", errDesc)
		}
		t.FailNow()
	}
}

func TestBookController_GetBookByISBN(t *testing.T) {
	bookController, tearDown := getBookController()
	defer tearDown()
	defer func(Collection db.BookCollection, isbn string) {
		_, _ = Collection.DeleteBookByISBN(isbn)
	}(bookController.Collection, testBook.ISBN)

	_ = bookController.Collection.CreateBook(testBook)

	req, _ := http.NewRequest(http.MethodGet, "/book/"+testBook.ISBN, nil)

	router := gin.Default()
	router.GET("/book/:isbn", bookController.GetBookByISBN)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Logf("Expected return code was %d, got %d", http.StatusOK, w.Code)

		var errDesc map[string]string
		if err := json.Unmarshal(w.Body.Bytes(), &errDesc); err != nil {
			t.Logf("Error with unmarshal: %s", err)
			t.FailNow()
		}

		if val, ok := errDesc["error"]; ok {
			t.Logf("Error description: %s", val)
		} else {
			t.Logf("Error description not found, response: %s", errDesc)
		}
		t.FailNow()
	}

	var book model.Book
	if err := json.Unmarshal(w.Body.Bytes(), &book); err != nil {
		t.Logf("Error with unmarshal: %s", err)
		t.FailNow()
	}

	if !testBook.Equal(book) {
		t.Log("Objects are not the same")
		t.Log("Expected:")
		t.Logf("\n%v", testBook)
		t.Log("Got:")
		t.Logf("\n%v", book)
		t.FailNow()
	}
}
