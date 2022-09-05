package service

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/korpgoodness/service.git/pkg/repository"
)

type SettingsService struct {
	repo repository.Settings
}

func NewSettingsService(repo repository.Settings) *SettingsService {
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

func (s *SettingsService) GetMarks(ctx context.Context) ([]domain.Mark, error) {
	marks, err := s.repo.GetMarks(ctx)
	return marks, err
}

func (s *SettingsService) SaveMarks(ctx context.Context, marks []domain.Mark) error {
	err := s.repo.SaveMarks(ctx, marks)
	return err
}

func (s *SettingsService) DeleteMark(ctx context.Context, mark domain.Mark) error {
	err := s.repo.DeleteMark(ctx, mark)
	return err
}
