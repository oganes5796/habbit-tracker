package service

import (
	"context"

	"github.com/oganes5796/habbit-tracker/internal/model"
	"github.com/oganes5796/habbit-tracker/internal/repository"
)

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *authService {
	return &authService{repo: repo}
}

func (s *authService) Create(ctx context.Context, user *model.AuthInfo) (int, error) {
	return s.repo.Create(ctx, user)
}
