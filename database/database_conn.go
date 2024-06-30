package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client *mongo.Client
}

func NewDatabaseClient(ctx context.Context) (*Database, error) {
	timeOutCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	dbURI := os.Getenv("MONGODB_URI")
	conn, err := mongo.Connect(timeOutCtx, options.Client().ApplyURI(dbURI))
	if err != nil {
		return nil, err
	}

	return &Database{client: conn}, nil
}

func (d *Database) Disconnect(ctx context.Context) {
	d.client.Disconnect(ctx)
}

func (d *Database) GetDatabase(databaseName string) *mongo.Database {
	return d.client.Database(databaseName, nil)
}
