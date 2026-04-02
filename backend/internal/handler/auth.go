package handler

import (
	"encoding/json"
	"net/http"

	"github.com/oganes5796/habbit-tracker/internal/model"
	"github.com/oganes5796/habbit-tracker/pkg/logger"
	"go.uber.org/zap"
)

func (im *Implemintation) CreateAuth(w http.ResponseWriter, r *http.Request) {
	var info model.AuthInfo

	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "internal error",
		})
		return
	}

	ctx := r.Context()
	id, err := im.serv.AuthService.Create(ctx, &info)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "error in creating auth handler",
		})
		logger.Error(ctx, "error handler create user", zap.Error(err))
		return
	}
	logger.Info(ctx, "Create user")

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"id": id,
	})
}
