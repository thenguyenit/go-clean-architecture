package entity

import "time"

//User represent the user entity
type User struct {
	ID          int64     `json:"id" bjon:"_id,omitempty"`
	Picture     string    `json:"picture" bjon:"picture,omitempty"`
	Email       string    `json:"email" bjon:"email"`
	Password    string    `json:"password" bjon:"password"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	ValidatedAt time.Time `json:"validated_at" bson:"validated_at,omitempty"`
}
