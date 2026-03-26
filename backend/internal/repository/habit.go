package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/oganes5796/habbit-tracker/internal/model"
)

const (
	tableName = "habits"
)

type habitRepo struct {
	conn *pgx.Conn
}

func NewHabitRepo(conn *pgx.Conn) *habitRepo {
	return &habitRepo{conn: conn}
}

func (r *habitRepo) Create(ctx context.Context, habit *model.HabitInfo) (int, error) {
	var id int
	query := "INSERT INTO " + tableName + " (user_id, title, type, target_value) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.conn.QueryRow(ctx, query, habit.UserID, habit.Title, habit.Type, habit.TargetValue).Scan(&id)
	if err != nil {
		return 0, errors.New("failed to create habit in repo: " + err.Error())
	}
	return id, nil
}
