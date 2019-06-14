package user

import (
	"context"

	"github.com/thenguyenit/go-clean-architecture/models"
)

type UserUsecase interface {
	Fetch(ctx context.Context, page int, number int) ([]*models.User, error)
	Insert(ctx context.Context, user *models.User) error
}
