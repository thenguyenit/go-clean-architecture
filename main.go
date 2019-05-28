package main

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo"
	config "github.com/micro/go-config"
	_userHttpDelivery "github.com/thenguyenit/go-clean-architecture/user/delivery/http"
	_userRepo "github.com/thenguyenit/go-clean-architecture/user/repository"
	_userUsecase "github.com/thenguyenit/go-clean-architecture/user/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AppConfig struct {
	Host      string              `json:"host"`
	Databases map[string]Database `json:"databases"`
}

type Database struct {
	URL     string `json:"url"`
	DBName  string `json:"database"`
	TimeOut int    `json:"timeout"`
}

var appConf *AppConfig

var dbConn *mongo.Database

func init() {

	//Load configuration
	if appConf == nil {
		config.LoadFile("./env.json")
		config.Scan(&appConf)
	}

	//Init DB Connection
	if dbConn == nil {

		mongoConf := appConf.Databases["mongo"]
		client, err := mongo.NewClient(options.Client().ApplyURI(mongoConf.URL))
		if err != nil {
			fmt.Println(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = client.Connect(ctx)
		if err != nil {
			fmt.Println(err)
		}
		dbConn = client.Database(mongoConf.DBName)
	}

}

func main() {

	e := echo.New()

	userRepo := _userRepo.NewUserRepository(dbConn)
	userUseCase := _userUsecase.NewUserUseCase(userRepo)
	_userHttpDelivery.NewUserHandler(e, userUseCase)

	e.Logger.Fatal(e.Start(appConf.Host))
}
