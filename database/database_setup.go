package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSetup() *mongo.Client {

	//load .env file
	err := godotenv.Load(".env")
	//check for error while loading .env
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("MONGO_URL")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	return client

}

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = OpenCollection(Client, collectionName)
	return collection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = OpenCollection(Client, collectionName)
	return collection
}

var Client *mongo.Client = DBSetup()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = (*mongo.Collection)(client.Database("E-Commerce").Collection(collectionName))
	return collection
}

var UserCollection *mongo.Collection = OpenCollection(Client, "Users")
