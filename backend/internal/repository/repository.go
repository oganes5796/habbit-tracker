package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/oganes5796/habbit-tracker/internal/model"
)

type HabitRepository interface {
	Create(ctx context.Context, habit *model.HabitInfo) (int, error)
}

type AuthRepository interface {
	Create(ctx context.Context, user *model.AuthInfo) (int, error)
}

type Repository struct {
	AuthRepository
	HabitRepository
}

func NewRepository(conn *pgx.Conn) *Repository {
	return &Repository{
		AuthRepository:  NewAuthRepo(conn),
		HabitRepository: NewHabitRepo(conn),
	}
}
