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

type Service struct {
	Authorization
	Folders
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		Folders:       NewFoldersService(repos),
	}
}
