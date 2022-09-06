package service

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/korpgoodness/service.git/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MarksService struct {
	repo repository.Marks
}

func NewMarksService(repo repository.Marks) *MarksService {
	return &MarksService{repo: repo}
}

func (s *MarksService) GetMarks(ctx context.Context) ([]domain.MarkGet, error) {
	marks, err := s.repo.GetMarks(ctx)
	return marks, err
}

func (s *MarksService) AddMark(ctx context.Context, mark domain.MarkCreate) error {
	err := s.repo.AddMark(ctx, mark)
	return err
}

func (s *MarksService) UpdateMark(ctx context.Context, markID primitive.ObjectID, mark domain.MarkCreate) error {
	err := s.repo.UpdateMark(ctx, markID, mark)
	return err
}

func (s *MarksService) DeleteMark(ctx context.Context, markID primitive.ObjectID) error {
	err := s.repo.DeleteMark(ctx, markID)
	return err
}
