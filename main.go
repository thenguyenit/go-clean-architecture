package main

import (
	"github.com/labstack/echo"
	config "github.com/micro/go-config"
	"github.com/thenguyenit/go-clean-architecture/models"
	_userHttpDelivery "github.com/thenguyenit/go-clean-architecture/user/delivery/http"
	_userRepo "github.com/thenguyenit/go-clean-architecture/user/repository"
	_userUsecase "github.com/thenguyenit/go-clean-architecture/user/usecase"
)

var appConf *models.AppConfig

func init() {

	//Load configuration
	if appConf == nil {
		config.LoadFile("./env.json")
		config.Scan(&appConf)
	}

}

func main() {

	e := echo.New()

	userRepo := _userRepo.NewUserRepository(appConf.Databases["mongo"])
	userUseCase := _userUsecase.NewUserUseCase(appConf.Usecases["user"], userRepo)
	_userHttpDelivery.NewUserHandler(e, userUseCase)

	e.Logger.Fatal(e.Start(appConf.Host))
}
