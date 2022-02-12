package service

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/korpgoodness/service.git/pkg/repository"
)

type UserDataService struct {
	repo repository.UserData
}

func NewUserDataService(repo repository.UserData) *UserDataService {
	return &UserDataService{repo: repo}
}

func (s *UserDataService) GetSettings(ctx context.Context) (domain.Settings, error) {
	settings, err := s.repo.GetSettings(ctx)
	return settings, err
}

func (s *UserDataService) SaveSettings(ctx context.Context, dataSettings domain.Settings) error {
	err := s.repo.SaveSettings(ctx, dataSettings)
	return err
}
