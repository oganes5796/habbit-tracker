package service

import (
	"context"

	"github.com/oganes5796/habbit-tracker/internal/model"
	"github.com/oganes5796/habbit-tracker/internal/repository"
)

type HabitService interface {
	Create(ctx context.Context, habit *model.HabitInfo) (int, error)
}

type Service struct {
	HabitService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		HabitService: NewHabitService(repo.HabitRepository),
	}
}
