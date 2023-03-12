package model

import "reflect"

type Book struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	ISBN string `json:"isbn" bson:"isbn" binding:"required"`

	Title       string   `json:"title" bson:"title" binding:"required"`
	Subtitle    string   `json:"subtitle" bson:"subtitle,omitempty"`
	Author      string   `json:"author" bson:"author" binding:"required"`
	Year        int      `json:"year" bson:"year" binding:"required,min=0"`
	Description string   `json:"description" bson:"description" binding:"required"`
	Categories  []string `json:"categories" bson:"categories"`

	OriginalTitle    string `json:"original_title" bson:"original_title,omitempty"`
	OriginalSubtitle string `json:"original_subtitle" bson:"original_subtitle,omitempty"`
	OriginalYear     int    `json:"original_year" bson:"original_year,omitempty" binding:"min=0"`
	Translator       string `json:"translator" bson:"translator,omitempty"`

	Size   string `json:"size" bson:"size" binding:"required"`
	Weight string `json:"weight" bson:"weight" binding:"required"`
	Pages  int    `json:"pages" bson:"pages" binding:"required,min=1"`

	Publisher string  `json:"publisher" bson:"publisher" binding:"required"`
	Language  string  `json:"language" bson:"language" binding:"required"`
	Price     float32 `json:"price" bson:"price" binding:"required,min=0"`
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
