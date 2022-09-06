package service

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/korpgoodness/service.git/pkg/repository"
)

type SettingsService struct {
	repo repository.SettingsInviting
}

func NewSettingsService(repo repository.SettingsInviting) *SettingsService {
	return &SettingsService{repo: repo}
}

func (s *SettingsService) GetSettings(ctx context.Context) (domain.Settings, error) {
	settings, err := s.repo.GetSettings(ctx)
	return settings, err
}

func (s *SettingsService) SaveSettings(ctx context.Context, dataSettings domain.Settings) error {
	err := s.repo.SaveSettings(ctx, dataSettings)
	return err
}
