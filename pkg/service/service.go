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

type SettingsInviting interface {
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
	CheckingUniqueness(ctx context.Context, phone string) (bool, error)
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
	CheckingUniqueness(ctx context.Context, channel_id string) (bool, error)
	Add(ctx context.Context, channel domain.ChannelAdd) error
	Get(ctx context.Context) ([]domain.ChannelGet, error)
	Launch(ctx context.Context, channelID primitive.ObjectID) error
	Update(ctx context.Context, channelID primitive.ObjectID, channel domain.ChannelIdKey) error
	Delete(ctx context.Context, channelID primitive.ObjectID, channel_id string) error
	EditChannel(ctx context.Context, channelID primitive.ObjectID, channel domain.CommentEdit) error
	EditProxy(ctx context.Context, channelID primitive.ObjectID, proxy string) error
	EditMark(ctx context.Context, channelID, mark primitive.ObjectID) error
}

type Marks interface {
	GetMarks(ctx context.Context) ([]domain.MarkGet, error)
	AddMark(ctx context.Context, mark domain.MarkCreate) error
	UpdateMark(ctx context.Context, markID primitive.ObjectID, mark domain.MarkCreate) error
	DeleteMark(ctx context.Context, markID primitive.ObjectID) error
}

// Structs

type AuthorizationService struct {
	Authorization
}

type InvitingService struct {
	SettingsInviting
	Folders
	Accounts
	AccountVerify
}

type AutomaticYoutubeService struct {
	Channels
	Marks
}

func NewAuthorizationService(repos *repository.AuthorizationRepository) *AuthorizationService {
	return &AuthorizationService{
		Authorization: NewAuthService(repos.Authorization),
	}
}

func NewInvitingService(repos *repository.InvitingRepository) *InvitingService {
	return &InvitingService{
		SettingsInviting: NewSettingsService(repos.SettingsInviting),
		Folders:          NewFoldersService(repos.Folders),
		Accounts:         NewAccountsService(repos.Accounts),
		AccountVerify:    NewAccountVerifyService(repos.Accounts),
	}
}

func NewAutomaticYoutubeService(repos *repository.AutomaticYoutubeRepository) *AutomaticYoutubeService {
	return &AutomaticYoutubeService{
		Channels: NewChannelsService(repos.Channels),
		Marks:    NewMarksService(repos.Marks),
	}
}
