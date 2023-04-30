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
//	@Failure		400			{object}	model.Error	"Bad request"
//	@Failure		500			{object}	model.Error	"Internal server error"
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

// GetBookByISBN godoc
//
//	@Summary		Get book by ISBN
//	@Description	Get book by ISBN
//	@Tags			books
//	@Param			isbn	path	string	true	"Book ISBN"
//	@Produce		json
//	@Success		200	{object}	model.Book	"Book object"
//	@Failure		404	{object}	model.Error	"Book not found"
//	@Failure		500	{object}	model.Error	"Internal server error"
//	@Router			/book/{isbn} [get]
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

// UpdateBook godoc
//
//	@Summary		Update book
//	@Description	Update book
//	@Tags			books
//	@Accept			json
//	@Produce		json
//	@Param			book	body	model.Book	true	"Book object"
//	@Success		204		"No content"
//	@Failure		400		{object}	model.Error	"Bad request"
//	@Failure		404		{object}	model.Error	"Book not found"
//	@Failure		500		{object}	model.Error	"Internal server error"
//	@Router			/book [put]
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

// DeleteBookByISBN godoc
//
//	@Summary		Delete book by ISBN
//	@Description	Delete book by ISBN
//	@Tags			books
//	@Param			isbn	path	string	true	"Book ISBN"
//	@Success		204		"No content"
//	@Failure		404		{object}	model.Error	"Book not found"
//	@Failure		500		{object}	model.Error	"Internal server error"
//	@Router			/book/{isbn} [delete]
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

// GetBooks godoc
//
//	@Summary		Get all books
//	@Description	Get all books
//	@Tags			books
//	@Produce		json
//	@Success		200	{object}	model.Book	"Book objects"
//	@Failure		404	{object}	model.Error	"Books not found"
//	@Failure		500	{object}	model.Error	"Internal server error"
//	@Router			/books [get]
func (receiver BookController) GetBooks(ctx *gin.Context) {
	books, err := receiver.Collection.GetBooks()

	if errors.Is(err, mongo.ErrNoDocuments) || len(books) == 0 {
		ctx.JSON(http.StatusNotFound, model.Error{Error: "books not found"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Error{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, books)
}
