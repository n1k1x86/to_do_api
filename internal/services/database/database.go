package database

import (
	"context"
	"database/sql"
	"log"
	"to_do/config"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"
)

func RunMigrations(cfg config.ToDoDB) error {
	db, err := sql.Open("pgx/v5", cfg.BuildDSN())
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return err
	}
	err = goose.Up(db, "migrations")
	if err != nil {
		return err
	}
	log.Println("Migrations were upped successfully")
	return nil
}

func CreatePool(ctx context.Context, cfg config.ToDoDB) (*pgxpool.Pool, error) {
	conf, err := pgxpool.ParseConfig(cfg.BuildDSN())
	if err != nil {
		return nil, err
	}
	conf.MaxConns = 10
	conf.MinConns = 2
	pool, err := pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
