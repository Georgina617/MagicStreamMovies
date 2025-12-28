package database

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Connect() (*mongo.Client, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: unable to find .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return nil, errors.New("MONGODB_URI not set")
	}

	log.Println("MONGODB_URI is set")

	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}

func OpenCollection(collectionName string, client *mongo.Client) *mongo.Collection {
	databaseName := os.Getenv("DATABASE_NAME")

	if databaseName == "" {
		log.Fatal("DATABASE_NAME not set!")
	}

	log.Println("DATABASE_NAME:", databaseName)

	collection := client.Database(databaseName).Collection(collectionName)

	if collection == nil {
		return nil
	}
	return collection
}
