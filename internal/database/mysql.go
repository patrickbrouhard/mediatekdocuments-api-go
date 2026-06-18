package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/patrickbrouhard/mediatekdocuments-api-go/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

func OpenMySQL(cfg config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("open mysql connection: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
