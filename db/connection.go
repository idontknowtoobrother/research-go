package db

import (
	"context"

	"github.com/research-mongo/collection/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ImplConnect() (*mongo.Client, error) {
	// Connect to the database
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		return client, err
	}

	// Check the connection

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return client, err
	}

	return client, nil
}
