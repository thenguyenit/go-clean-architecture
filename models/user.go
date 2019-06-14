package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User represent the user entity
type User struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// ID          int64              `json:"id"`
	Username    string    `json:"username" bson:"username,omitempty"`
	Email       string    `json:"email" bson:"email"`
	Password    string    `json:"password" bson:"password"`
	Status      int       `json:"status" bson:"status"`
	Roles       string    `json:"roles" bson:"roles"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	ValidatedAt time.Time `json:"validated_at" bson:"validated_at,omitempty"`
}
