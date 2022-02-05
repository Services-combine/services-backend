package service

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/korpgoodness/service.git/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDataService struct {
	repo repository.UserData
}

func NewUserDataService(repo repository.UserData) *UserDataService {
	return &UserDataService{repo: repo}
}

func (s *UserDataService) GetSettings(ctx context.Context, userID primitive.ObjectID) (domain.Settings, error) {
	settings, err := s.repo.GetSettings(ctx, userID)
	return settings, err
}

func (s *UserDataService) SaveSettings(ctx context.Context, userID primitive.ObjectID, dataSettings domain.Settings) error {
	err := s.repo.SaveSettings(ctx, userID, dataSettings)
	return err
}
