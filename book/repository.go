package book

//Repository represent the book repository
type Repository interface {
	Find(id int64) (*Book, error)
	// FindAll() ([]*Book, error)
	// Add(book *Book) (int64, error)
	// Update(book *Book) error
	// Delete(id int64) error
}
