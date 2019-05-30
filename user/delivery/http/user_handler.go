package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/thenguyenit/go-clean-architecture/models"

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
	e.POST("/users", handler.InsertUser)
}

func (u *UserHandler) FetchUser(c echo.Context) error {
	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)
	pageS := c.QueryParam("page")
	if pageS == "" {
		pageS = "1"
	}
	page, _ := strconv.Atoi(pageS)

	ctx := c.Request().Context()

	listAr, _, err := u.UserUsecase.Fetch(ctx, int64(page), int64(num))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}
	// c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, listAr)
}

func (u *UserHandler) InsertUser(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = u.UserUsecase.Insert(ctx, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)

}
