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

type Folders interface {
	Get(ctx context.Context, path string) ([]domain.Folder, error)
	Create(ctx context.Context, folder domain.Folder) error
	GetData(ctx context.Context, hash string) (domain.Folder, error)
	Move(ctx context.Context, hash, path string) error
	Rename(ctx context.Context, hash, name string) error
	ChangeChat(ctx context.Context, hash, chat string) error
	ChangeUsernames(ctx context.Context, hash string, usernames []string) error
	ChangeMessage(ctx context.Context, hash, message string) error
	ChangeGroups(ctx context.Context, hash string, groups []string) error
	Delete(ctx context.Context, hash string) error
}

type Repository struct {
	Authorization
	Folders
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Authorization: NewAuthRepo(db.Database(viper.GetString("mongo.databaseName"))),
		Folders:       NewFoldersRepo(db.Database(viper.GetString("mongo.databaseName"))),
	}
}
