package service

import (
	"context"
	domain_settings "github.com/b0shka/services/internal/domain/settings"
	"github.com/b0shka/services/internal/repository"
)

type SettingsService struct {
	repo repository.Settings
}

func NewSettingsService(repo repository.Settings) *SettingsService {
	return &SettingsService{repo: repo}
}

func (s *SettingsService) Get(ctx context.Context) (domain_settings.Settings, error) {
	settings, err := s.repo.Get(ctx, repository.ServiceInviting)
	return settings, err
}

func (s *SettingsService) Save(ctx context.Context, inp domain_settings.SaveSettingsInput) error {
	settingsParams := repository.SaveSettingsParams{
		Service:       repository.ServiceInviting,
		CountInviting: inp.CountInviting,
		CountMailing:  inp.CountMailing,
	}
	err := s.repo.Save(ctx, settingsParams)
	return err
}
