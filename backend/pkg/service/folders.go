package service

import (
	"context"
	"math/rand"
	"time"

	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/korpgoodness/service.git/pkg/repository"
)

type FoldersService struct {
	repo repository.Folders
}

func NewFoldersService(repo repository.Folders) *FoldersService {
	return &FoldersService{repo: repo}
}

func GenerateHash() string {
	const LENGTH_HASH = 34
	const symbols = "1234567890qwertyuiopasdfghjklzxcvbnm"
	random_hash := make([]byte, LENGTH_HASH)

	rand.Seed(time.Now().UnixNano())
	for i := range random_hash {
		random_hash[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(random_hash)
}

func (s *FoldersService) Get(ctx context.Context, path string) ([]domain.Folder, error) {
	folders, err := s.repo.Get(ctx, path)
	return folders, err
}

func (s *FoldersService) Create(ctx context.Context, folder domain.Folder) error {
	hash := GenerateHash()
	folder.Hash = hash

	err := s.repo.Create(ctx, folder)
	return err
}

func (s *FoldersService) GetData(ctx context.Context, hash string) (domain.Folder, error) {
	folder, err := s.repo.GetData(ctx, hash)
	return folder, err
}

func (s *FoldersService) Move(ctx context.Context, hash, path string) error {
	err := s.repo.Move(ctx, hash, path)
	return err
}

func (s *FoldersService) Rename(ctx context.Context, hash, name string) error {
	err := s.repo.Rename(ctx, hash, name)
	return err
}

func (s *FoldersService) ChangeChat(ctx context.Context, hash, chat string) error {
	err := s.repo.ChangeChat(ctx, hash, chat)
	return err
}

func (s *FoldersService) ChangeUsernames(ctx context.Context, hash string, usernames []string) error {
	err := s.repo.ChangeUsernames(ctx, hash, usernames)
	return err
}

func (s *FoldersService) ChangeGroups(ctx context.Context, hash string, groups []string) error {
	err := s.repo.ChangeGroups(ctx, hash, groups)
	return err
}

func (s *FoldersService) Delete(ctx context.Context, hash string) error {
	err := s.repo.Delete(ctx, hash)
	return err
}
