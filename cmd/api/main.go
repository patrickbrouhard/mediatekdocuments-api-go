// cmd/api/main.go
package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/patrickbrouhard/mediatekdocuments-api-go/internal/config"
	apihttp "github.com/patrickbrouhard/mediatekdocuments-api-go/internal/http"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg := config.Load()
	router := apihttp.NewRouter()

	logger.Info("starting server",
		"addr", cfg.HTTPAddr,
		"env", cfg.AppEnv,
	)

	if err := http.ListenAndServe(cfg.HTTPAddr, router); err != nil {
		logger.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
