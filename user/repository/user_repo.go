package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/thenguyenit/go-clean-architecture/models"
	"github.com/thenguyenit/go-clean-architecture/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const USER_COLLECTION = "user"

type userrepo struct {
	dbConn *mongo.Database
}

func NewUserRepository(dbConn *mongo.Database) user.UserRepository {
	return &userrepo{dbConn}
}

func (r *userrepo) Fetch(ctx context.Context, page int, limit int) ([]*models.User, error) {

	//Query
	collection := r.dbConn.Collection(USER_COLLECTION)
	lm := int64(limit)
	skip := (int64(page) - 1) * lm

	findOptions := &options.FindOptions{Limit: &lm, Skip: &skip}

	cur, err := collection.Find(context.TODO(), bson.D{}, findOptions)
	defer cur.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	result := make([]*models.User, 0)
	for cur.Next(context.TODO()) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, &user)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *userrepo) Insert(ctx context.Context, user *models.User) error {
	insertResult, err := r.dbConn.Collection(USER_COLLECTION).InsertOne(ctx, user)
	if err != nil {
		return err
	}

	// user.ID = insertResult.InsertedID.(primitive.ObjectID)

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return nil
}
