package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	timeout = 10 * time.Second
)

type Config struct {
	Username string
	Password string
	DBName   string
}

func NewMongoDB(cfg Config) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.ykyjv.mongodb.net/%s?retryWrites=true&w=majority", cfg.Username, cfg.Password, cfg.DBName)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
