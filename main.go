package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	Id        int `json:"id" bson:"_id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

var collection *mongo.Collection

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading env file", err)	
	}

	// Connecting to DB
	mongoUri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)	
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)	
	}
	fmt.Println("Welcome to MongoDB Atlas")
}