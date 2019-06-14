package usecase

import (
	"context"

	"github.com/thenguyenit/go-clean-architecture/models"
	"github.com/thenguyenit/go-clean-architecture/user"
)

type userUsecase struct {
	userRepo user.UserRepository
}

func NewUserUseCase(ur user.UserRepository) user.UserUsecase {
	return &userUsecase{userRepo: ur}
}

func (u *userUsecase) Fetch(ctx context.Context, page int, number int) ([]*models.User, error) {
	if number == 0 {
		number = 10
	}

	listItem, err := u.userRepo.Fetch(ctx, page, number)

	if err != nil {
		return nil, err
	}
	return listItem, nil
}

func (u *userUsecase) Insert(ctx context.Context, user *models.User) error {
	err := u.userRepo.Insert(ctx, user)
	return err
}
