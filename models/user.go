package models

import "time"

//User represent the user entity
type User struct {
	ID          int64     `json:"id" bson:"_id,omitempty"`
	Username    string    `json:"username" bson:"username,omitempty"`
	Email       string    `json:"email" bson:"email"`
	Password    string    `json:"password" bson:"password"`
	Status      int       `json:"status" bson:"status"`
	Roles       string    `json:"roles" bson:"roles"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	ValidatedAt time.Time `json:"validated_at" bson:"validated_at,omitempty"`
}
