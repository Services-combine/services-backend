package service

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/korpgoodness/service.git/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Authorization

type Authorization interface {
	Login(ctx context.Context, username, password string) (domain.UserDataAuth, error)
	ParseToken(token string) (string, error)
	CheckUser(ctx context.Context, userID primitive.ObjectID) (domain.UserReduxData, error)
}

// Inviting

type Settings interface {
	GetSettings(ctx context.Context) (domain.Settings, error)
	SaveSettings(ctx context.Context, dataSettings domain.Settings) error
}

type Folders interface {
	GetFolders(ctx context.Context) (map[string]interface{}, error)
	Create(ctx context.Context, folder domain.Folder) error
	GetAllDataFolderById(ctx context.Context, folderID primitive.ObjectID) (map[string]interface{}, error)
	GetFoldersMove(ctx context.Context, folderID primitive.ObjectID) ([]domain.AccountDataMove, error)
	GetFoldersByPath(ctx context.Context, path string) ([]domain.FolderItem, error)
	GetFolderById(ctx context.Context, folderID primitive.ObjectID) (domain.Folder, error)
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
	Update(ctx context.Context, account domain.AccountUpdate) error
	Delete(ctx context.Context, accountID primitive.ObjectID) error
	GenerateInterval(ctx context.Context, folderID primitive.ObjectID) error
	CheckBlock(ctx context.Context, folderID primitive.ObjectID) error
}

type AccountVerify interface {
	LoginApi(ctx context.Context, accountID primitive.ObjectID) error
	ParsingApi(ctx context.Context, accountLogin domain.AccountLogin) error
	GetCodeSession(ctx context.Context, accountID primitive.ObjectID) error
	CreateSession(ctx context.Context, accountLogin domain.AccountLogin) error
}

// AutomaticYoutube

type Channels interface {
	Add(ctx context.Context, channel domain.ChannelAdd) error
}

// Structs

type AuthorizationService struct {
	Authorization
}

type InvitingService struct {
	Settings
	Folders
	Accounts
	AccountVerify
}

type AutomaticYoutubeService struct {
	Channels
}

func NewAuthorizationService(repos *repository.AuthorizationRepository) *AuthorizationService {
	return &AuthorizationService{
		Authorization: NewAuthService(repos.Authorization),
	}
}

func NewInvitingService(repos *repository.InvitingRepository) *InvitingService {
	return &InvitingService{
		Settings:      NewSettingsService(repos.Settings),
		Folders:       NewFoldersService(repos.Folders),
		Accounts:      NewAccountsService(repos.Accounts),
		AccountVerify: NewAccountVerifyService(repos.Accounts),
	}
}

func NewAutomaticYoutubeService(repos *repository.AutomaticYoutubeRepository) *AutomaticYoutubeService {
	return &AutomaticYoutubeService{
		Channels: NewChannelsService(repos.Channels),
	}
}
