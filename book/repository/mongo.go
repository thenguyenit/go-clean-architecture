package repository

import "go.mongodb.org/mongo-driver/mongo"

type repo struct {
	pool *mgosession.Pool
}

func NewRepository(dbConn *mongo.Database) Repository {

}

func (r *repo) Find(id int64) (*Book, error) {
	book := Book{
		ID:      1,
		Title:   "Sapiens",
		Content: "The story talk about the history of the person",
	}

	return &book, nil
}
