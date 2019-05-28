package book

import "github.com/thenguyenit/go-clean-architecture/catalog-service/entity"

//Repository represent the book repository
type Repository interface {
	Find(id int64) (*entity.Book, error)
	FindAll() ([]*entity.Book, error)
	Add(book *entity.Book) (int64, error)
	Update(book *entity.Book) error
	Delete(id int64) error
}

type repo struct {
	pool *mgosession.Pool
}

func (r *repo) Find(id int64) (*entity.Book, error) {
	book := entity.Book{
		ID:      1,
		Title:   "Sapiens",
		Content: "The story talk about the history of the person",
	}

	return &book, nil
}
