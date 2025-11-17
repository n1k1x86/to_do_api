package database

import (
	"context"
	"database/sql"
	"fmt"
	"todoapi/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
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

func NewPool(ctx context.Context, cfg *config.Database) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(BuildDSN(cfg))
	if err != nil {
		return nil, err
	}
	config.MaxConns = 10
	config.MinConns = 2

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
