package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

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

	http.ListenAndServe(port, handler.New())
}
