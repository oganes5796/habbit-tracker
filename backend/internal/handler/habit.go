package handler

import (
	"encoding/json"
	"net/http"

	"github.com/oganes5796/habbit-tracker/internal/model"
)

func (im *Implemintation) Create(w http.ResponseWriter, r *http.Request) {
	var info model.HabitInfo

	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "internal error",
		})
		return
	}

	id, err := im.serv.HabitService.Create(r.Context(), &info)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "error in creating habit handler",
		})

		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"id": id,
	})
}
