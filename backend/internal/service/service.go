package service

import (
	"context"

	"github.com/oganes5796/habbit-tracker/internal/model"
	"github.com/oganes5796/habbit-tracker/internal/repository"
)

type HabitService interface {
	Create(ctx context.Context, habit *model.HabitInfo) (int, error)
}

type AuthService interface {
	Create(ctx context.Context, user *model.AuthInfo) (int, error)
}

type Service struct {
	AuthService
	HabitService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AuthService:  NewAuthService(repo.AuthRepository),
		HabitService: NewHabitService(repo.HabitRepository),
	}
}
