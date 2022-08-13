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
	CheckUser(ctx context.Context, userID primitive.ObjectID) (domain.UserReduxData, error)
}

type Folders interface {
	GetCountAllAccount(ctx context.Context) (domain.AccountsCount, error)
	GetListFolders(ctx context.Context, path string) ([]domain.FolderItem, error)
	Get(ctx context.Context, path string) ([]domain.Folder, error)
	Create(ctx context.Context, folder domain.Folder) error
	GetData(ctx context.Context, folderID primitive.ObjectID) (domain.Folder, error)
	GetFolders(ctx context.Context) ([]domain.Folder, error)
	GetAccountByFolderID(ctx context.Context, folderID primitive.ObjectID, limitFolder domain.LimitFolder) ([]domain.Account, error)
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
	GetSettings(ctx context.Context) (domain.Settings, error)
}

type Accounts interface {
	Create(ctx context.Context, accountCreate domain.Account) error
	GetSettings(ctx context.Context, accountID primitive.ObjectID) (domain.AccountSettings, error)
	GetData(ctx context.Context, accountID primitive.ObjectID) (domain.Account, error)
	GetAccountsFolder(ctx context.Context, folderID primitive.ObjectID) ([]domain.Account, error)
	GetFolders(ctx context.Context) (map[string]string, error)
	GetFolderByID(ctx context.Context, folderID primitive.ObjectID) (domain.Folder, error)
	UpdateAccount(ctx context.Context, account domain.AccountUpdate) error
	Delete(ctx context.Context, accountID primitive.ObjectID) error
	GenerateInterval(ctx context.Context, folderID primitive.ObjectID) error
	AddRandomHash(ctx context.Context, accountID primitive.ObjectID, randomHash string) error
	AddPhoneHash(ctx context.Context, accountID primitive.ObjectID, phoneCodeHash string) error
	AddApi(ctx context.Context, accountSettings domain.AccountApi) error
	ChangeStatusBlock(ctx context.Context, accountID primitive.ObjectID, status string) error
	ChangeVerify(ctx context.Context, accountID primitive.ObjectID) error
}

type UserData interface {
	GetSettings(ctx context.Context) (domain.Settings, error)
	SaveSettings(ctx context.Context, dataSettings domain.Settings) error
}

type Repository struct {
	Authorization
	Folders
	Accounts
	UserData
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Authorization: NewAuthRepo(db.Database(viper.GetString("mongo.databaseName"))),
		Folders:       NewFoldersRepo(db.Database(viper.GetString("mongo.databaseName"))),
		Accounts:      NewAccountsRepo(db.Database(viper.GetString("mongo.databaseName"))),
		UserData:      NewUserDataRepo(db.Database(viper.GetString("mongo.databaseName"))),
	}
}
