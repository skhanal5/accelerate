package main

import (
	"accelerate/internal/handler"
	"errors"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	port := "8080" // should be a const

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.GetHealth)
	
	logger.Info("Starting the server", "port", port) //this is an antipattern?
	err := http.ListenAndServe(port, mux)

	if errors.Is(err, http.ErrServerClosed) {
		logger.Error(err.Error()) //this is also an antipattern
	}
}
