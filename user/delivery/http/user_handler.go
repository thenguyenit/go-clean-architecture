package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/thenguyenit/go-clean-architecture/user"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserUsecase user.Usecase
}

func NewUserHandler(e *echo.Echo, us user.Usecase) {
	handler := &UserHandler{us}

	e.GET("/users", handler.FetchUser)
}

func (u *UserHandler) FetchUser(c echo.Context) error {
	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)
	cursor := c.QueryParam("cursor")

	ctx := c.Request().Context()

	listAr, _, err := u.UserUsecase.Fetch(ctx, cursor, int64(num))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}
	// c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, listAr)
}
