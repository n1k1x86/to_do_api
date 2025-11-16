package database

import (
	"database/sql"
	"fmt"
	"todoapi/internal/config"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func BuildDSN(cfg *config.Database) string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.Username, cfg.Password, cfg.Hostname, cfg.DBName)
}

func UpMigrations(cfg *config.Database) error {
	db, err := sql.Open("postgres", BuildDSN(cfg))
	if err != nil {
		return err
	}
	defer db.Close()
	if err = goose.SetDialect("postgres"); err != nil {
		return err
	}
	if err = goose.Up(db, cfg.MigrationsDir); err != nil {
		return err
	}
	return nil
}
