package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/thenguyenit/go-clean-architecture/connections"
	"github.com/thenguyenit/go-clean-architecture/models"
	"github.com/thenguyenit/go-clean-architecture/user"
	"go.mongodb.org/mongo-driver/bson"
)

const USER_COLLECTION = "user"

type repo struct {
	dbConf models.DatabaseConf
}

func NewUserRepository(dbConf models.DatabaseConf) user.Repository {
	return &repo{dbConf}
}

func (r *repo) Fetch(ctx context.Context, cursor string, number int64) ([]*models.User, string, error) {

	//Make a connection to database
	dbClient, _ := connections.GetMongoInstance(ctx, r.dbConf)

	//Query
	collection := dbClient.Database(r.dbConf.DBName).Collection(USER_COLLECTION)
	cur, err := collection.Find(context.Background(), bson.D{})
	defer cur.Close(context.Background())
	defer dbClient.Disconnect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	result := make([]*models.User, 0)
	for cur.Next(ctx) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result....
		fmt.Println(result)
		result = append(result, &user)
	}

	if err := cur.Err(); err != nil {
		return nil, "", err
		log.Fatal(err)
	}

	return result, "", nil
}
