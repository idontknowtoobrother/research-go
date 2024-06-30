package main

import (
	"context"
	"os"
	"time"

	logCharmBracelet "github.com/charmbracelet/log"
	"github.com/idontknowtoobrother/mongo-models/database"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	log = logCharmBracelet.NewWithOptions(os.Stderr, logCharmBracelet.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "HEX 👻",
	})
)

var (
	credentialsCollection *mongo.Collection
	usersCollection       *mongo.Collection
)

func main() {

	ctx := context.Background()

	databaseClient, err := database.NewDatabaseClient(ctx)
	if err != nil {
		log.Fatal("Failed to connect to database", "error", err)
	}
	defer databaseClient.Disconnect(ctx)

	database := databaseClient.GetDatabase(os.Getenv("DB_NAME"))

	usersCollection := database.Collection(os.Getenv("COLLECTION_USERS"))
	credentialsCollection := database.Collection(os.Getenv("COLLECTION_CREDENTIALS"))

	if credentialsCollection == nil || usersCollection == nil {
		log.Fatal("Failed to get collections")
	}

	res, err := usersCollection.InsertOne(ctx, bson.D{
		{Key: "username", Value: "admin"},
		{Key: "password", Value: "pass"},
	})
	if err != nil {
		log.Fatal("Failed to insert document", "error", err)
	}

	log.Info("Inserted document", "id", res.InsertedID)

	log.Info("Connected to database")

}
