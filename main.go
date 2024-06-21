package main

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DBClient          *mongo.Client
	bookingCollection *mongo.Collection
	ocppCollection    *mongo.Collection
)

func LookDuplicateCollectionCall(wg *sync.WaitGroup) {
	defer wg.Done()

	mockBooking := bson.M{
		"booking_id": "123",
		"status":     "active",
	}

	mockOcpp := bson.M{
		"ocpp_id": "123",
		"status":  "active",
	}

	_, err := bookingCollection.InsertOne(context.Background(), mockBooking)
	if err != nil {
		log.Fatal("Error inserting booking:", err)
	}

	_, err = ocppCollection.InsertOne(context.Background(), mockOcpp)
	if err != nil {
		log.Fatal("Error inserting ocpp:", err)
	}

	// update booking
	_, err = bookingCollection.UpdateOne(context.Background(), bson.M{"booking_id": "123"}, bson.M{"$set": bson.M{"status": "inactive"}})
	if err != nil {
		log.Fatal("Error updating booking:", err)
	}

	// update ocpp
	_, err = ocppCollection.UpdateOne(context.Background(), bson.M{"ocpp_id": "123"}, bson.M{"$set": bson.M{"status": "inactive"}})
	if err != nil {
		log.Fatal("Error updating ocpp:", err)
	}

}
func main() {
	var err error
	DBClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer func() {
		if err = DBClient.Disconnect(context.TODO()); err != nil {
			log.Fatal("Error disconnecting from the database:", err)
		}
	}()

	bookingCollection = DBClient.Database("historic").Collection("booking")
	ocppCollection = DBClient.Database("historic").Collection("ocpp")

	var wg sync.WaitGroup
	numWorkers := 100 // จำนวนโกRoutine ที่ทำงานพร้อมกัน
	jobQueue := make(chan struct{}, numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			for range jobQueue {
				LookDuplicateCollectionCall(&wg)
			}
		}()
	}

	for i := 0; i < 500000; i++ {
		wg.Add(1)
		jobQueue <- struct{}{}
	}

	close(jobQueue)
	wg.Wait()
}
