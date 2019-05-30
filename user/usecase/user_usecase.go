package usecase

import (
	"context"

	"github.com/thenguyenit/go-clean-architecture/models"
	"github.com/thenguyenit/go-clean-architecture/user"
)

type userUsecase struct {
	userRepo user.Repository
}

func NewUserUseCase(ur user.Repository) user.Usecase {
	return &userUsecase{userRepo: ur}
}

func (u *userUsecase) Fetch(ctx context.Context, page int64, number int64) ([]*models.User, string, error) {
	if number == 0 {
		number = 10
	}

	listItem, err := u.userRepo.Fetch(ctx, page, number)

	if err != nil {
		return nil, "", err
	}
	return listItem, "", nil
}

func (u *userUsecase) Insert(ctx context.Context, user *models.User) error {
	err := u.userRepo.Insert(ctx, user)
	return err
}
