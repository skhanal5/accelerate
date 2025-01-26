package main

import (
	"accelerate/internal/handler"
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
	logger.Error(http.ListenAndServe(":8080", mux).Error())
}
