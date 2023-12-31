package main

import (
	"log/slog"
	"net/http"

	"github.com/kyong0612/fitness-saporter/handler"
)

func main() {
	slog.Info("Server is running on port 8080")

	http.ListenAndServe(":8080", handler.New())
}
