package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Функция для подключения к базе данных
func ConnectToDB() error {
	uri := "mongodb://admin:Bo5aK5!t@212.67.11.16:27017/tama"
	clientOptions := options.Client().ApplyURI(uri).SetAuth(options.Credential{
		Username:      "admin",
		Password:      "Bo5aK5!t",
		AuthSource:    "admin",
		AuthMechanism: "SCRAM-SHA-256",
	})

	var err error
	client, err = mongo.NewClient(clientOptions)
	if err != nil {
		return fmt.Errorf("ошибка создания клиента: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return fmt.Errorf("ошибка подключения: %v", err)
	}

	return nil
}

// Функция для получения коллекции пользователей
func GetUserCollection() *mongo.Collection {
	return client.Database("tama").Collection("ta-ma-db")
}
