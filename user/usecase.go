package user

import (
	"context"

	"github.com/thenguyenit/go-clean-architecture/models"
)

type Usecase interface {
	Fetch(ctx context.Context, page int64, number int64) ([]*models.User, string, error)
	Insert(ctx context.Context, user *models.User) error
}
