package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"main/db"
	"main/model"
	"net/http"
	"strings"
)

type BookController struct {
	Collection db.BookCollection
}

// CreateBook godoc
//
//	@Summary		Create new book
//	@Description	Create new book
//	@Tags			books
//	@Accept			json
//	@Produce		json
//	@Param			requestBody	body	model.Book	true	"Book object"
//	@Success		201			"No content"
//	@Failure		400			{object}	model.Error	"error: bad request"
//	@Failure		500			{object}	model.Error	"error: internal server error"
//	@Router			/book [post]
func (receiver BookController) CreateBook(ctx *gin.Context) {
	var book model.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	err := receiver.Collection.CreateBook(book)
	if err != nil {
		if strings.Contains(err.Error(), "book with isbn="+book.ISBN+" already exists") {
			ctx.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, model.Error{Error: err.Error()})
		return
	}
	ctx.Status(http.StatusCreated)
}

func (receiver BookController) GetBookByISBN(ctx *gin.Context) {
	book, err := receiver.Collection.GetBookByISBN(ctx.Param("isbn"))

	if errors.Is(err, mongo.ErrNoDocuments) {
		ctx.JSON(http.StatusNotFound, model.Error{Error: "document with isbn=" + ctx.Param("isbn") + " not found"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Error{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func (receiver BookController) UpdateBook(ctx *gin.Context) {
	var book model.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Error{Error: err.Error()})
		return
	}

	updated, err := receiver.Collection.UpdateBook(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Error{Error: err.Error()})
		return
	}

	if updated == 0 {
		ctx.JSON(http.StatusNotFound, model.Error{Error: "document with isbn=" + book.ISBN + " not found"})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (receiver BookController) DeleteBookByISBN(ctx *gin.Context) {
	deleted, err := receiver.Collection.DeleteBookByISBN(ctx.Param("isbn"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Error{Error: err.Error()})
		return
	}

	if deleted == 0 {
		ctx.JSON(http.StatusNotFound, model.Error{Error: "document with isbn=" + ctx.Param("isbn") + " not found"})
		return
	}
	ctx.Status(http.StatusNoContent)
}
