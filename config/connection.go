package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func MongoConnect() *mongo.Client {
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(LoadEnv("MONGO_URI")))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	
	return client

}

var Connection *mongo.Client = MongoConnect()

func GetCollection(client *mongo.Client , collection string) *mongo.Collection{
	return client.Database(LoadEnv("DB_NAME")).Collection(collection)
}