package model

type Book struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	ISBN string `json:"isbn" bson:"isbn"`

	Title       string   `json:"title" bson:"title"`
	Subtitle    string   `json:"subtitle" bson:"subtitle,omitempty"`
	Author      string   `json:"author" bson:"author"`
	Year        int      `json:"year" bson:"year"`
	Description string   `json:"description" bson:"description"`
	Categories  []string `json:"categories" bson:"categories"`

	OriginalTitle    string `json:"original_title" bson:"original_title,omitempty"`
	OriginalSubtitle string `json:"original_subtitle" bson:"original_subtitle,omitempty"`
	OriginalYear     int    `json:"original_year" bson:"original_year,omitempty"`
	Translator       string `json:"translator" bson:"translator,omitempty"`

	Size   string `json:"size" bson:"size"`
	Weight string `json:"weight" bson:"weight"`
	Pages  int    `json:"pages" bson:"pages"`

	Publisher string  `json:"publisher" bson:"publisher"`
	Language  string  `json:"language" bson:"language"`
	Price     float32 `json:"price" bson:"price"`
}
