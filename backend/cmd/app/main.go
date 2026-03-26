package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/oganes5796/habbit-tracker/internal/client"
	"github.com/oganes5796/habbit-tracker/internal/config"
	"github.com/oganes5796/habbit-tracker/internal/handler"
	"github.com/oganes5796/habbit-tracker/internal/repository"
	"github.com/oganes5796/habbit-tracker/internal/server"
	"github.com/oganes5796/habbit-tracker/internal/service"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	cfg := config.NewCfgDB()
	pool, err := client.NewPostgresDB(ctx, cfg)
	if err != nil {
		panic("Failed to connect to the database")
	}
	defer pool.Close(ctx)
	slog.Info("Successfully connected to PostgreSQL")

	repository := repository.NewRepository(pool)
	service := service.NewService(repository)
	handlers := handler.NewImplementation(service)

	srv := &server.Server{}
	go func() {
		if err := srv.Run(
			os.Getenv("HOST"),
			os.Getenv("PORT"),
			handlers.InitRoutes(),
		); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("error occurred while running http server", "error", err)
		}
	}()
	slog.Info("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("App shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("error occurred on server shutting down", "error", err)
	}

	slog.Info("App exited")
}
