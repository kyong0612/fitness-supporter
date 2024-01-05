package handler

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type handler struct{}

func New() http.Handler {
	r := chi.NewRouter()
	h := handler{} // NOTE: inject dependencies here

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/healthcheck"))

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("I'm alive!"))
		if err != nil {
			slog.Error(err.Error())
		}
	})

	r.Group(func(r chi.Router) {
		r.Post("/line/webhook", h.PostLINEWebhook)
	})

	return r
}
