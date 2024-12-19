package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type User struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

func InitMongoDB(uri string) {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
}

func FindUserByEmail(email string) (*User, error) {
	collection := client.Database("tama").Collection("ta-ma-db")
	var user User
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		log.Printf("Error finding user: %v", err)
		return nil, err
	}
	log.Printf("Found user: %+v", user)
	return &user, nil
}
