package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/kyong0612/fitness-saporter/handler"
	"github.com/kyong0612/fitness-saporter/infra/config"
)

func main() {
	// Load config
	if err := config.New(); err != nil {
		slog.Error(err.Error())
		os.Exit(1) // Exit with error.
	}

	port := fmt.Sprintf(":%d", config.Get().Port)
	slog.Info("Server is running on port " + port)

	http.ListenAndServe(port, handler.New())
}
