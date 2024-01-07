package handler

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kyong0612/fitness-supporter/generated/proto/handler/v1/analyzeimagev1connect"
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

	path, analyzeImageHandler := analyzeimagev1connect.NewAnalyzeImageServiceHandler(h)

	// connect-go
	r.Route(path, func(r chi.Router) {
		r.Post("/AnalyzeImage", analyzeImageHandler.ServeHTTP)
	})

	r.Group(func(r chi.Router) {
		r.Post("/line/webhook", h.PostLINEWebhook)
	})

	return r
}
