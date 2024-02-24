package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/kyong0612/fitness-supporter/proto/generated/proto/handler/v1/handlerv1connect"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
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

	r.Use(otelhttp.NewMiddleware("handler"))

	// pubsub endpoints
	analyzeImagePath, analyzeImageHandler := handlerv1connect.NewAnalyzeImageServiceHandler(h)
	RMUAppleHealthcarePath, RMUAppleHealthcareHandler := handlerv1connect.NewRMUAppleHealthcareServiceHandler(h)

	fmt.Println("analyzeImagePath", analyzeImagePath)
	fmt.Println("RMUAppleHealthcarePath", RMUAppleHealthcarePath)

	r.Group(func(r chi.Router) {
		r.Post(analyzeImagePath+"AnalyzeImage", analyzeImageHandler.ServeHTTP)
		r.Post(RMUAppleHealthcarePath+"RMUAppleHealthcare", RMUAppleHealthcareHandler.ServeHTTP)
	})

	r.Group(func(r chi.Router) {
		r.Post("/line/webhook", h.PostLINEWebhook)
	})

	r.Group(func(r chi.Router) {
		r.Post("/sync/healthcare/apple", h.SyncHealthcareApple)
	})

	return r
}
