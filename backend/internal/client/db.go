package client

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/oganes5796/habbit-tracker/internal/config"
)

func NewPostgresDB(ctx context.Context, cfg config.Config) (*pgx.Conn, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SSLMode)

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		slog.Error("NewPostgresDB:Connect", "error", err)
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		conn.Close(ctx)
		slog.Error("NewPostgresDB:Ping", "error", err)
		return nil, err
	}

	return conn, nil
}
