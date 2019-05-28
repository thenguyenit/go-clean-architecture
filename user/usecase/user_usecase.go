package usecase

import (
	"github.com/thenguyenit/go-clean-architecture/models"
	"github.com/thenguyenit/go-clean-architecture/user"
)

type userUsecase struct {
	userRepo user.Repository
}

func NewUserUseCase(ur user.Repository) user.Usecase {
	return &userUsecase{ur}
}

func (u *userUsecase) Fetch() ([]*models.User, error) {
	listUser, err := u.userRepo.Fetch()
	if err != nil {
		return nil, err
	}
	return listUser, nil
}
