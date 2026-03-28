package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/oganes5796/habbit-tracker/internal/model"
)

const (
	tableUsersName = "users"
)

type authRepo struct {
	conn *pgx.Conn
}

func NewAuthRepo(conn *pgx.Conn) *authRepo {
	return &authRepo{conn: conn}
}

func (r *authRepo) Create(ctx context.Context, user *model.AuthInfo) (int, error) {
	var id int
	err := r.conn.QueryRow(ctx, "INSERT INTO "+tableUsersName+" (username) VALUES ($1) RETURNING id", user.Username).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
