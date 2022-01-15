package service

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/korpgoodness/service.git/pkg/repository"
)

type Authorization interface {
	GenerateToken(ctx context.Context, username, password string) (string, error)
	ParseToken(token string) (string, error)
}

type Inviting interface {
	GetFolders(ctx context.Context, path string) ([]domain.Folder, error)
	CreateFolder(ctx context.Context, folder domain.Folder) error
	GetDataFolder(ctx context.Context, hash string) (domain.Folder, error)
	RenameFolder(ctx context.Context, hash, name string) error
}

type Service struct {
	Authorization
	Inviting
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		Inviting:      NewInvitingService(repos),
	}
}
