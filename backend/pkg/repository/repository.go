package repository

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	GetUser(ctx context.Context, username, password string) (domain.User, error)
}

type Inviting interface {
	GetFolders(ctx context.Context, path string) ([]domain.Folder, error)
	CreateFolder(ctx context.Context, folder domain.Folder) error
	GetDataFolder(ctx context.Context, hash string) (domain.Folder, error)
	RenameFolder(ctx context.Context, hash, name string) error
}

type Repository struct {
	Authorization
	Inviting
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Authorization: NewAuthMongoDB(db.Database(viper.GetString("mongo.databaseName"))),
		Inviting:      NewInvitingMongoDB(db.Database(viper.GetString("mongo.databaseName"))),
	}
}
