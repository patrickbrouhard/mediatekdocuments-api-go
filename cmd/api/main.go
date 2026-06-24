// cmd/api/main.go
package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/patrickbrouhard/mediatekdocuments-api-go/internal/config"
	"github.com/patrickbrouhard/mediatekdocuments-api-go/internal/database"
	apihttp "github.com/patrickbrouhard/mediatekdocuments-api-go/internal/http"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := config.Load()

	db, err := database.OpenMySQL(cfg)
	if err != nil {
		logger.Error("failed to open database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	router := apihttp.NewRouter(db, logger)

	logger.Info("starting server",
		"addr", cfg.HTTPAddr,
		"env", cfg.AppEnv,
		"db_host", cfg.DBHost,
		"db_name", cfg.DBName,
	)

	if err := http.ListenAndServe(cfg.HTTPAddr, router); err != nil {
		logger.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
