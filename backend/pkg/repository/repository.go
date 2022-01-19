package repository

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	GetUser(ctx context.Context, username, password string) (domain.User, error)
}

type Folders interface {
	Get(ctx context.Context, path string) ([]domain.Folder, error)
	Create(ctx context.Context, folder domain.Folder) error
	GetData(ctx context.Context, folderID primitive.ObjectID) (domain.Folder, error)
	GetAccountByFolderID(ctx context.Context, folderID primitive.ObjectID) ([]domain.Account, error)
	GetCountAccounts(ctx context.Context, folderID primitive.ObjectID) (domain.AccountsCount, error)
	Move(ctx context.Context, folderID primitive.ObjectID, path string) error
	Rename(ctx context.Context, folderID primitive.ObjectID, name string) error
	ChangeChat(ctx context.Context, folderID primitive.ObjectID, chat string) error
	ChangeUsernames(ctx context.Context, folderID primitive.ObjectID, usernames []string) error
	ChangeMessage(ctx context.Context, folderID primitive.ObjectID, message string) error
	ChangeGroups(ctx context.Context, folderID primitive.ObjectID, groups []string) error
	Delete(ctx context.Context, folderID primitive.ObjectID) error
	LaunchInviting(ctx context.Context, folderID primitive.ObjectID) error
	LaunchMailingUsernames(ctx context.Context, folderID primitive.ObjectID) error
	LaunchMailingGroups(ctx context.Context, folderID primitive.ObjectID) error
}

type Accounts interface {
	Create(ctx context.Context, accountCreate domain.Account) error
	GetData(ctx context.Context, accountID primitive.ObjectID) (domain.Account, error)
	GetFolderByID(ctx context.Context, folderID primitive.ObjectID) (domain.Folder, error)
	UpdateAccount(ctx context.Context, account domain.AccountUpdate) error
	Delete(ctx context.Context, accountID primitive.ObjectID) error
	GenerateInterval(ctx context.Context, folderID primitive.ObjectID) error
}

type Repository struct {
	Authorization
	Folders
	Accounts
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Authorization: NewAuthRepo(db.Database(viper.GetString("mongo.databaseName"))),
		Folders:       NewFoldersRepo(db.Database(viper.GetString("mongo.databaseName"))),
		Accounts:      NewAccountsRepo(db.Database(viper.GetString("mongo.databaseName"))),
	}
}
