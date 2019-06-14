package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Book represent the book entity
type Book struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string             `json:"title" validate:"required"`
	Content   string             `json:"content" validate:"required"`
	Author    Author             `json:"author"`
	UpdatedAt time.Time          `json:"updated_at"`
	CreatedAt time.Time          `json:"created_at"`
}
