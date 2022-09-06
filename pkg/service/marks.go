package service

import (
	"context"

	"github.com/korpgoodness/service.git/internal/domain"
	"github.com/korpgoodness/service.git/pkg/repository"
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

func (s *MarksService) UpdateMark(ctx context.Context, mark domain.MarkGet) error {
	err := s.repo.UpdateMark(ctx, mark)
	return err
}
