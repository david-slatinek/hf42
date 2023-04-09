package model

import "reflect"

// Book is a model for book
//
//	@Description	Book model
type Book struct {
	// ISBN is the unique identifier of the book
	ISBN string `json:"isbn" bson:"isbn" binding:"required" example:"978-3-16-148410-0" validate:"required" minLength:"17" maxLength:"17"`

	// Title is the title of the book
	Title string `json:"title" bson:"title" binding:"required" example:"The Hitchhiker's Guide to the Galaxy" validate:"required" minLength:"1"`
	// Subtitle is the subtitle of the book
	Subtitle string `json:"subtitle" bson:"subtitle,omitempty" example:"A Trilogy in Five Parts" validate:"optional" minLength:"1"`
	// Author is the author of the book
	Author string `json:"author" bson:"author" binding:"required" example:"Douglas Adams" validate:"required" minLength:"1"`
	// Year is the year of publication
	Year int `json:"year" bson:"year" binding:"required,min=0" example:"1979" validate:"required" min:"0"`
	// Description is a short description of the book
	Description string `json:"description" bson:"description" binding:"required" example:"Go on a galactic adventure with the last human on Earth, his alien best friend, and a depressed android." validate:"required" minLength:"1"`
	// Categories is a list of categories
	Categories []string `json:"categories" bson:"categories" example:"'Science Fiction', 'Fantasy'" validate:"required" minLength:"1"`

	// OriginalTitle is the title of the original book
	OriginalTitle string `json:"originalTitle" bson:"original_title,omitempty" example:"The Hitchhiker's Guide to the Galaxy" validate:"optional" minLength:"1"`
	// OriginalSubtitle is the subtitle of the original book
	OriginalSubtitle string `json:"originalSubtitle" bson:"original_subtitle,omitempty" example:"A Trilogy in Five Parts" validate:"optional" minLength:"1"`
	// OriginalAuthor is the author of the original book
	OriginalYear int `json:"originalYear" bson:"original_year,omitempty" binding:"min=0" example:"1979" validate:"optional" min:"0"`
	// OriginalDescription is a short description of the original book
	Translator string `json:"translator" bson:"translator,omitempty" example:"John Stone" validate:"optional" minLength:"1"`

	// Size is the size of the book
	Size string `json:"size" bson:"size" binding:"required" example:"21 x 14 cm" validate:"required" minLength:"1"`
	// Weight is the weight of the book
	Weight string `json:"weight" bson:"weight" binding:"required" example:"0.3 kg" validate:"required" minLength:"1"`
	// Pages is the number of pages in the book
	Pages int `json:"pages" bson:"pages" binding:"required,min=1" example:"215" validate:"required" min:"1"`

	// Publisher is the publisher of the book
	Publisher string `json:"publisher" bson:"publisher" binding:"required" example:"Pan Books Ltd" validate:"required" minLength:"1"`
	// Language is the language of the book
	Language string `json:"language" bson:"language" binding:"required" example:"English" validate:"required" minLength:"1"`

	// Quantity amount of books
	Quantity int `json:"quantity" binding:"required" validate:"required" min:"1"`
	// Price is the price of the book
	Price float32 `json:"price" bson:"price" binding:"required,min=0" example:"21.99" validate:"required" min:"0"`

	// TotalPrice total amount for book, Quantity * Price
	TotalPrice float32 `json:"totalPrice" binding:"required" validate:"required" min:"0"`
}

func (receiver Book) Equal(book Book) bool {
	if len(receiver.Categories) != len(book.Categories) {
		return false
	}

	for i := range receiver.Categories {
		if receiver.Categories[i] != book.Categories[i] {
			return false
		}
	}

	val := reflect.ValueOf(&receiver).Elem()
	otherFields := reflect.Indirect(reflect.ValueOf(book))

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		if typeField.Name == "ID" || typeField.Name == "Categories" {
			continue
		}

		value := val.Field(i)
		otherValue := otherFields.FieldByName(typeField.Name)

		if value.Interface() != otherValue.Interface() {
			return false
		}
	}
	return true
}
