package entity

import "time"

//Author represent the author entity
type Author struct {
	ID        int64     `json:"id" bjon:"_id"`
	Name      string    `json:"string" bjon:"string"`
	CreatedAt time.Time `json:"created_at" bjon:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bjon:"updated_at"`
}
