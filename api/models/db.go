package models

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Collection

// initialize connection to mongo db
func init() {
	_ = godotenv.Load()
	dbURL, ok := os.LookupEnv("DB_URL")
	if !ok {
		log.Fatal("db env not set")
	}
	clientOptions := options.Client().ApplyURI(dbURL)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the db...")
	db = client.Database("go-testing").Collection("users")
}
