package controller

import (
	"github.com/gin-gonic/gin"
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

	id, err := receiver.Collection.Insert(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}
