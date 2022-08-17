package service

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/korpgoodness/service.git/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userData struct {
	AccessToken  string
	RefreshToken string
	UserID       string
}

// Services

type Authorization interface {
	Login(ctx context.Context, username, password string) (userData, error)
	ParseToken(token string) (string, error)
	CheckUser(ctx context.Context, userID primitive.ObjectID) (domain.UserReduxData, error)
}


// Inviting 

type Folders interface {
	GetDataMainPage(ctx context.Context) (map[string]interface{}, error) //
	GetListFolders(ctx context.Context, path string) ([]domain.FolderItem, error) //
	Get(ctx context.Context, path string) ([]domain.Folder, error) //
	Create(ctx context.Context, folder domain.Folder) error
	GetData(ctx context.Context, folderID primitive.ObjectID) (domain.Folder, error) //
	GetFolderById(ctx context.Context, folderID primitive.ObjectID) (map[string]interface{}, error)
	GetFoldersMove(ctx context.Context, folderID primitive.ObjectID) ([]domain.DataFolderHash, error) //
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
	UpdateAccount(ctx context.Context, account domain.AccountUpdate) error
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

type UserData interface {
	GetSettings(ctx context.Context) (domain.Settings, error)
	SaveSettings(ctx context.Context, dataSettings domain.Settings) error
}


// Channels

//...


type AuthorizetionService struct {
	Authorization
}

type InvitingService struct {
	Folders
	Accounts
	AccountVerify
	UserData
}

type ChannelsService struct {

}


func NewAuthorizetionService(repos *repository.Repository) *AuthorizetionService {
	return &AuthorizetionService{
		Authorization: NewAuthService(repos.Authorization),
	}
}

func NewInviting(repos *repository.Repository) *InvitingService {
	return &InvitingService{
		Folders:       NewFoldersService(repos.Folders),
		Accounts:      NewAccountsService(repos.Accounts),
		AccountVerify: NewAccountVerifyService(repos.Accounts),
		UserData:      NewUserDataService(repos.UserData),
	}
}

func NewChannels(repos *repository.Repository) *ChannelsService {
	return &ChannelsService{
		
	}
}