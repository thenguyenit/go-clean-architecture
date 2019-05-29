package usecase

import (
	"context"

	"github.com/thenguyenit/go-clean-architecture/models"
	"github.com/thenguyenit/go-clean-architecture/user"
)

type userUsecase struct {
	config   models.UsecaseConf
	userRepo user.Repository
}

func NewUserUseCase(conf models.UsecaseConf, ur user.Repository) user.Usecase {
	return &userUsecase{config: conf, userRepo: ur}
}

func (u *userUsecase) Fetch(ctx context.Context, cursor string, number int64) ([]*models.User, string, error) {
	if number == 0 {
		number = 10
	}

	listItem, cursor, err := u.userRepo.Fetch(ctx, cursor, number)

	if err != nil {
		return nil, "", err
	}
	return listItem, "", nil
}
