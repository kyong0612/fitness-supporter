package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/kyong0612/fitness-saporter/config"
	"github.com/kyong0612/fitness-saporter/handler"
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
