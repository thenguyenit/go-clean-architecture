package connections

import (
	"context"

	"github.com/thenguyenit/go-clean-architecture/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoInstance(ctx context.Context, dbConf models.DatabaseConf) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(dbConf.URL))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}
