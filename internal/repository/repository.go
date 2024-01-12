package repository

import (
	"context"
	"github.com/b0shka/services/internal/domain"
	domain_auth "github.com/b0shka/services/internal/domain/auth"
	domain_settings "github.com/b0shka/services/internal/domain/settings"
	domain_user "github.com/b0shka/services/internal/domain/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users interface {
	Get(ctx context.Context, arg GetUserParams) (domain_user.User, error)
	Check(ctx context.Context, id primitive.ObjectID) (domain_user.UserReduxData, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) error
	GetSession(ctx context.Context, id primitive.ObjectID) (domain_auth.Session, error)
}

type Settings interface {
	Get(ctx context.Context, service string) (domain_settings.Settings, error)
	Save(ctx context.Context, arg SaveSettingsParams) error
}

type Folders interface {
	GetFolders(ctx context.Context) ([]domain.Folder, error)
	GetFoldersByPath(ctx context.Context, path string) ([]domain.FolderItem, error)
	Create(ctx context.Context, folder domain.Folder) error
	GetFolderById(ctx context.Context, folderID primitive.ObjectID) (domain.Folder, error)
	GetAccountsByFolderID(ctx context.Context, folderID primitive.ObjectID) ([]domain.Account, error)
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
	CheckingUniqueness(ctx context.Context, phone string) (bool, error)
	Create(ctx context.Context, accountCreate domain.Account) error
	Update(ctx context.Context, account domain.AccountUpdate) error
	Delete(ctx context.Context, accountID primitive.ObjectID) error
	GetById(ctx context.Context, accountID primitive.ObjectID) (domain.Account, error)
	GetAccountsByFolderID(ctx context.Context, folderID primitive.ObjectID) ([]domain.Account, error)
	GenerateInterval(ctx context.Context, folderID primitive.ObjectID) error
	ChangeStatusBlock(ctx context.Context, accountID primitive.ObjectID, status string) error
	GetGroupById(ctx context.Context, folderID primitive.ObjectID) (string, error)
}

type Repositories struct {
	Users
	Settings
	Folders
	Accounts
}

func NewRepositories(db *mongo.Client, databaseName string) *Repositories {
	return &Repositories{
		Users:    NewUsersRepo(db.Database(databaseName)),
		Settings: NewSettingsRepo(db.Database(databaseName)),
		Folders:  NewFoldersRepo(db.Database(databaseName)),
		Accounts: NewAccountsRepo(db.Database(databaseName)),
	}
}
