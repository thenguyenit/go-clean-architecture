package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/thenguyenit/go-clean-architecture/models"
	"github.com/thenguyenit/go-clean-architecture/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	DBConn         *mongo.Database
	CollectionName string
}

func NewUserRepository(dbConn *mongo.Database) user.Repository {
	return &repo{dbConn, "user"}
}

func (r *repo) Fetch() ([]*models.User, error) {
	collection := r.DBConn.Collection(r.CollectionName)
	ctx := context.Background()
	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

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
		return nil, err
		log.Fatal(err)
	}

	return result, nil
}
