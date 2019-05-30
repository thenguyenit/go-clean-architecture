package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
	"github.com/thenguyenit/go-clean-architecture/user/mocks"
)

func TestFetch(t *testing.T) {
	mockUserRepo := new(mocks.Repository)
	assert.Equal(t, 1, 1, "This shoud be equal")

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Fetch", mock.Anything, mock.AnythingOfType("int64"))
	})
}
