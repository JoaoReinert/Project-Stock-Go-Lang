package main

import (
	"log/slog"
)

func main() {
	slog.Info("Starting server, hello!")

	err := Start()
	if err != nil {
		slog.Error("Failed to start the server", slog.String("error", err.Error()))
	}
}
