package user

import (
	"context"

	"github.com/thenguyenit/go-clean-architecture/models"
)

//UserRepository represent the book repository
type UserRepository interface {
	Fetch(ctx context.Context, page int, number int) ([]*models.User, error)
	Insert(ctx context.Context, user *models.User) error
	// FindAll() ([]*Book, error)
	// Add(book *Book) (int64, error)
	// Update(book *Book) error
	// Delete(id int64) error
}
