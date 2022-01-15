package service

import (
	"context"
	"fmt"
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
	folders, err := s.repo.GetFolders(ctx, "/")
	if err != nil {
		return nil, err
	}

	return folders, err
}

func (s *InvitingService) CreateFolder(ctx context.Context, folder domain.Folder) error {
	hash := GenerateHash()
	fmt.Println(hash)
	folder.Hash = hash
	if err := s.repo.CreateFolder(ctx, folder); err != nil {
		return err
	}
	return nil
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
