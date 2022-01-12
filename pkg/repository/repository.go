package repository

import (
	"context"

	combine "github.com/korpgoodness/services.git"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	GetUser(ctx context.Context, username, password string) (combine.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Authorization: NewAuthMongoDB(db.Database(viper.GetString("mongo.databaseName")).Collection(usersCollection)),
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
