package database

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var Client *mongo.Client
var Ctx context.Context

func Connect(coll string) *mongo.Collection {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("mongoURI")))
	Client = client
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	Ctx = ctx

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	demoDB := client.Database("hotel")
	collection := demoDB.Collection(coll)
	return collection

}
