package service

import (
	"context"
	"math/rand"

	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/korpgoodness/service.git/pkg/repository"
)

type InvitingService struct {
	repo repository.Inviting
}

func NewInvitingService(repo repository.Inviting) *InvitingService {
	return &InvitingService{repo: repo}
}

func (s *InvitingService) GetFolders(ctx context.Context, path string) ([]domain.Folder, error) {
	folders, err := s.repo.GetFolders(ctx, path)
	return folders, err
}

func (s *InvitingService) CreateFolder(ctx context.Context, folder domain.Folder) error {
	hash := GenerateHash()
	folder.Hash = hash

	err := s.repo.CreateFolder(ctx, folder)
	return err
}

func GenerateHash() string {
	LENGTH_HASH := 34
	symbols := []rune("1234567890qwertyuiopasdfghjklzxcvbnm")

	random_hash := make([]rune, LENGTH_HASH)
	for i := range random_hash {
		random_hash[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(random_hash)
}

func (s *InvitingService) GetDataFolder(ctx context.Context, hash string) (domain.Folder, error) {
	folder, err := s.repo.GetDataFolder(ctx, hash)
	return folder, err
}

func (s *InvitingService) RenameFolder(ctx context.Context, hash, name string) error {
	err := s.repo.RenameFolder(ctx, hash, name)
	return err
}
