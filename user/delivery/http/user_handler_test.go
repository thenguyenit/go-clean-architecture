package http_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bxcodec/faker"
	"github.com/labstack/echo"
	"github.com/thenguyenit/go-clean-architecture/models"
	userHttp "github.com/thenguyenit/go-clean-architecture/user/delivery/http"
	"github.com/thenguyenit/go-clean-architecture/user/mocks"
)

func TestFetchUser(t *testing.T) {

	userFaked := models.User{
		Username: faker.Username(),
		Email:    faker.Email(),
	}
	mockListUser := make([]*models.User, 0)
	mockListUser = append(mockListUser, &userFaked)

	mockUserUsecase := new(mocks.UserUsecase)
	mockUserUsecase.On("Fetch", context.Background(), 1, 0).Return(mockListUser, nil)

	//Mock echo
	e := echo.New()
	req := httptest.NewRequest("Get", "/users", strings.NewReader(""))

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("users")

	userHandler := userHttp.UserHandler{
		UserUsecase: mockUserUsecase,
	}

	userHandler.FetchUser(c)
	jB, _ := json.Marshal(mockListUser)
	assert.Equal(t, string(jB), strings.Replace(rec.Body.String(), "\n", "", 1))

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUserUsecase.AssertCalled(t, "Fetch", context.Background(), 1, 0)
}
