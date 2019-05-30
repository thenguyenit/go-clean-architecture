package main

import (
	"context"
	"log"

	"github.com/labstack/echo"
	config "github.com/micro/go-config"
	"github.com/thenguyenit/go-clean-architecture/connections"
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

	//Init mongo db connection
	dbClient, err := connections.GetMongoInstance(context.Background(), appConf.Databases["mongo"])
	if err != nil {
		log.Fatal(err)
	}
	defer dbClient.Disconnect(context.Background())
	mDBConn := dbClient.Database(appConf.Databases["mongo"].DBName)

	userRepo := _userRepo.NewUserRepository(mDBConn)
	userUseCase := _userUsecase.NewUserUseCase(userRepo)
	_userHttpDelivery.NewUserHandler(e, userUseCase)

	e.Logger.Fatal(e.Start(appConf.Host))
}
