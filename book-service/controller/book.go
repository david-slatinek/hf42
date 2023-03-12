package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"main/db"
	"main/model"
	"net/http"
)

type BookController struct {
	Collection db.BookCollection
}

func (receiver BookController) CreateBook(ctx *gin.Context) {
	var book model.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.ID = ""

	id, err := receiver.Collection.Insert(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (receiver BookController) GetBookByISBN(ctx *gin.Context) {
	book, err := receiver.Collection.GetBookByISBN(ctx.Param("isbn"))

	if errors.Is(err, mongo.ErrNoDocuments) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "document with isbn=" + ctx.Param("isbn") + " not found"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (receiver BookController) UpdateBook(ctx *gin.Context) {
	var book model.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.ID = ""

	updated, err := receiver.Collection.UpdateBook(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if updated == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "document with isbn=" + book.ISBN + " not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"updated": updated})
}
