package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/charoleizer/my-collection/src/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func loadEnv() error {
	err := config.LoadDatabaseEnv()
	if err != nil {
		return err
	}

	return nil
}

func clientOptions() *options.ClientOptions {
	loadEnv()

	user := os.Getenv("mongodb_username")
	pass := os.Getenv("momgodb_password")
	url := os.Getenv("mongodb_url")

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	return options.Client().
		ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s", user, pass, url)).
		SetServerAPIOptions(serverAPIOptions)
}

func getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func Connect() (*mongo.Client, error) {
	clientOptions := clientOptions()

	ctx, cancel := getContext()
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	return client, nil
}
