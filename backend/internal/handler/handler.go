package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/oganes5796/habbit-tracker/internal/service"
)

type Implemintation struct {
	serv *service.Service
}

func NewImplementation(serv *service.Service) *Implemintation {
	return &Implemintation{serv: serv}
}

func (im *Implemintation) InitRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"hello"}`))
	})

	r.Route("/api", func(r chi.Router) {
		r.Post("/", im.Create)
	})

	return r
}
