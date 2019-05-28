package user

import "github.com/thenguyenit/go-clean-architecture/models"

type Usecase interface {
	Fetch() ([]*models.User, error)
}
