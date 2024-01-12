package service

import (
	"context"
	"github.com/b0shka/services/internal/config"
	"github.com/b0shka/services/internal/domain"
	domain_auth "github.com/b0shka/services/internal/domain/auth"
	domain_folders "github.com/b0shka/services/internal/domain/folders"
	domain_settings "github.com/b0shka/services/internal/domain/settings"
	"github.com/b0shka/services/internal/repository"
	"github.com/b0shka/services/internal/worker"
	"github.com/b0shka/services/pkg/auth"
	"github.com/b0shka/services/pkg/hash"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Auth interface {
	Login(ctx *gin.Context, inp domain_auth.LoginInput) (domain_auth.LoginOutput, error)
	RefreshToken(ctx context.Context, inp domain_auth.RefreshTokenInput) (domain_auth.RefreshTokenOutput, error)
}

type Settings interface {
	Get(ctx context.Context) (domain_settings.Settings, error)
	Save(ctx context.Context, inp domain_settings.SaveSettingsInput) error
}

type Folders interface {
	GetFolders(ctx context.Context) (domain_folders.GetFoldersOutput, error)
	Create(ctx context.Context, folder domain.Folder) error
	GetAllDataFolderById(ctx context.Context, folderID primitive.ObjectID) (domain_folders.GetFolderOutput, error)
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
	JoinGroup(ctx context.Context, folderID primitive.ObjectID) error
}

type Services struct {
	Auth
	Settings
	Folders
	Accounts
}

type Deps struct {
	Repos        *repository.Repositories
	Hasher       hash.Hasher
	TokenManager auth.Manager
	//OTPGenerator    otp.Generator
	//IDGenerator     identity.Generator
	AuthConfig      config.AuthConfig
	TaskDistributor worker.TaskDistributor
}

func NewServices(deps Deps) *Services {
	return &Services{
		Auth: NewAuthService(
			deps.Repos.Users,
			deps.Hasher,
			deps.TokenManager,
			deps.AuthConfig,
			deps.TaskDistributor,
		),
		Settings: NewSettingsService(deps.Repos.Settings),
		Folders: NewFoldersService(
			deps.Repos.Folders,
			deps.Repos.Settings,
		),
		Accounts: NewAccountsService(deps.Repos.Accounts),
	}
}
