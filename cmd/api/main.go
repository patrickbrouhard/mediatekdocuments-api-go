package main

import (
	"log/slog"
	"net/http"
	"os"

	apihttp "github.com/patrickbrouhard/mediatekdocuments-api-go/internal/http"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	router := apihttp.NewRouter()

	logger.Info("Starting server", "addr", ":8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.Error("Server stopped", "error", err)
		os.Exit(1)
	}
}
