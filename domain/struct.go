package domain

import "time"

type Nota struct {
	ID           bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	IdNota       int           `json:"id_nota" bson:"id_nota"`
	Category     string        `json:"category" bson:"category"`
	CategoryName string        `json:"category_name" bson:"category_name"`
	Deporte      []string      `json:"deporte" bson:"deporte"`
	DeporteName  []string      `json:"deporte_name" bson:"deporte_name"`
	Type         string        `json:"type" bson:"type"`
	Title        string        `json:"title" bson:"title"`
	Description  string        `json:"description" bson:"description"`
	Slug         string        `json:"slug" bson:"slug"`
	Link         string        `json:"link" bson:"link"`
	Image        string        `json:"image" bson:"image"`
	Content      string        `json:"content" bson:"content"`
	Pais         []string      `json:"pais" bson:"pais"`
	PaisName     []string      `json:"pais_name" bson:"pais_name"`
	Author       string        `json:"author" bson:"author"`
	PublishDate  time.Time     `json:"publish_date" bson:"publish_date"`
	ModifiedDate time.Time     `json:"modified_date" bson:"modified_date"`
	Status       string        `json:"status" bson:"status"`
}

