package user

import (
	"context"

	"github.com/thenguyenit/go-clean-architecture/models"
)

type Usecase interface {
	Fetch(ctx context.Context, cursor string, number int64) ([]*models.User, string, error)
}
