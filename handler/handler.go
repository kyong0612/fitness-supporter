package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func New() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I'm alive!"))
	})

	return r
}
