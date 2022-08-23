package repository

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Authorization

type Authorization interface {
	GetUser(ctx context.Context, username, password string) (domain.User, error)
	CheckUser(ctx context.Context, userID primitive.ObjectID) (domain.UserReduxData, error)
}

// Inviting

type Settings interface {
	GetSettings(ctx context.Context) (domain.Settings, error)
	SaveSettings(ctx context.Context, dataSettings domain.Settings) error
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
	GetSettings(ctx context.Context) (domain.Settings, error)
}

type Accounts interface {
	Create(ctx context.Context, accountCreate domain.Account) error
	Update(ctx context.Context, account domain.AccountUpdate) error
	Delete(ctx context.Context, accountID primitive.ObjectID) error
	GetById(ctx context.Context, accountID primitive.ObjectID) (domain.Account, error)
	GetAccountsByFolderID(ctx context.Context, folderID primitive.ObjectID) ([]domain.Account, error)
	GenerateInterval(ctx context.Context, folderID primitive.ObjectID) error
	AddRandomHash(ctx context.Context, accountID primitive.ObjectID, randomHash string) error
	AddPhoneHash(ctx context.Context, accountID primitive.ObjectID, phoneCodeHash string) error
	AddApi(ctx context.Context, accountSettings domain.AccountApi) error
	ChangeStatusBlock(ctx context.Context, accountID primitive.ObjectID, status string) error
	ChangeVerify(ctx context.Context, accountID primitive.ObjectID) error
}

// AutomaticYoutube

type Channels interface {
	CheckingUniqueness(ctx context.Context, channel_id string) (bool, error)
	Add(ctx context.Context, channel domain.ChannelAdd) error
	Get(ctx context.Context) ([]domain.ChannelGet, error)
	Launch(ctx context.Context, channelID primitive.ObjectID) error
	Update(ctx context.Context, channelID primitive.ObjectID, channel domain.ChannelUpdate) error
	Delete(ctx context.Context, channelID primitive.ObjectID) error
	Edit(ctx context.Context, channelID primitive.ObjectID, channel domain.ChannelEdit) error
}

type AuthorizationRepository struct {
	Authorization
}

type InvitingRepository struct {
	Folders
	Accounts
	Settings
}

type AutomaticYoutubeRepository struct {
	Channels
}

func NewAuthRepository(db *mongo.Client) *AuthorizationRepository {
	return &AuthorizationRepository{
		Authorization: NewAuthRepo(db.Database(viper.GetString("mongo.databaseName"))),
	}
}

func NewInvitingRepository(db *mongo.Client) *InvitingRepository {
	return &InvitingRepository{
		Folders:  NewFoldersRepo(db.Database(viper.GetString("mongo.databaseName"))),
		Accounts: NewAccountsRepo(db.Database(viper.GetString("mongo.databaseName"))),
		Settings: NewUserDataRepo(db.Database(viper.GetString("mongo.databaseName"))),
	}
}

func NewAutomaticYoutubeRepository(db *mongo.Client) *AutomaticYoutubeRepository {
	return &AutomaticYoutubeRepository{
		Channels: NewChannelsRepo(db.Database(viper.GetString("mongo.databaseName"))),
	}
}
