package service

import (
	"context"

	"github.com/oganes5796/habbit-tracker/internal/model"
	"github.com/oganes5796/habbit-tracker/internal/repository"
)

type habitService struct {
	repo repository.HabitRepository
}

func NewHabitService(repo repository.HabitRepository) *habitService {
	return &habitService{repo: repo}
}

func (s *habitService) Create(ctx context.Context, habit *model.HabitInfo) (int, error) {
	return s.repo.Create(ctx, habit)
}
