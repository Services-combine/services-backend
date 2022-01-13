package repository

import (
	"context"

	"github.com/korpgoodness/services.git/internal/domain"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	GetUser(ctx context.Context, username, password string) (domain.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Authorization: NewAuthMongoDB(db.Database(viper.GetString("mongo.databaseName"))),
	}
}
