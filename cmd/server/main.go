package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/kyong0612/fitness-saporter/handler"
	"github.com/kyong0612/fitness-saporter/infra/config"
	"github.com/kyong0612/fitness-saporter/infra/logger"
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

	srv := &http.Server{
		Addr:              port,
		Handler:           handler.New(),
		ReadHeaderTimeout: 20 * time.Second,
		WriteTimeout:      20 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		slog.Error(err.Error())
		os.Exit(1) // Exit with error.
	}
}
