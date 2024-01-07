package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/kyong0612/fitness-supporter/handler"
	"github.com/kyong0612/fitness-supporter/infra/config"
	"github.com/kyong0612/fitness-supporter/infra/logger"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	// Load config
	if err := config.New(); err != nil {
		slog.Error(err.Error())
		os.Exit(1) // Exit with error.
	}

	// Init logger
	logger.Init()

	port := fmt.Sprintf(":%d", config.Get().Port)
	slog.Info("Server is running on port " + port)

	r := handler.New()

	h2s := &http2.Server{}
	srv := &http.Server{
		Addr:              port,
		ReadHeaderTimeout: 20 * time.Second,
		WriteTimeout:      20 * time.Second,
		Handler:           h2c.NewHandler(r, h2s), // HTTP/2 Cleartext handler
	}

	if err := srv.ListenAndServe(); err != nil {
		slog.Error(err.Error())
		os.Exit(1) // Exit with error.
	}
}
