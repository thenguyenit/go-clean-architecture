package usecase_test

import (
	"context"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/thenguyenit/go-clean-architecture/models"

	"github.com/thenguyenit/go-clean-architecture/user/mocks"
	"github.com/thenguyenit/go-clean-architecture/user/usecase"
)

func TestFetch(t *testing.T) {

	mockUserRepo := new(mocks.UserRepository)
	mockListUser := make([]*models.User, 0)

	userFaked := models.User{
		Username: faker.Username(),
		Email:    faker.Email(),
	}
	mockListUser = append(mockListUser, &userFaked)

	mockUserRepo.On("Fetch", context.TODO(), 1, 10).Return(mockListUser, nil)

	userUsecase := usecase.NewUserUseCase(mockUserRepo)
	list, err := userUsecase.Fetch(context.TODO(), 1, 10)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(list))

	mockUserRepo.AssertCalled(t, "Fetch", context.TODO(), mock.AnythingOfType("int"), mock.AnythingOfType("int"))

}

func TestInsert(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)

	userFaked := models.User{
		Username: faker.Username(),
		Email:    faker.Email(),
	}

	mockUserRepo.On("Insert", context.TODO(), &userFaked).Return(nil)
	userUsecase := usecase.NewUserUseCase(mockUserRepo)

	err := userUsecase.Insert(context.TODO(), &userFaked)

	assert.NoError(t, err)
	mockUserRepo.AssertCalled(t, "Insert", context.TODO(), &userFaked)
}
