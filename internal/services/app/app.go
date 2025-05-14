package app

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Ctx        context.Context
	CancelFunc context.CancelFunc
	Pool       *pgxpool.Pool
}

func (a *App) Run() {
	recover()
}

func (a *App) Close() {
	if a.Pool != nil {
		log.Println("Closing connections with the database...")
		a.Pool.Close()
	}
	log.Println("Closing the application...")
	a.CancelFunc()
}

func CreateNewApp(ctx context.Context, cancelFunc context.CancelFunc, pool *pgxpool.Pool) *App {
	return &App{
		Pool:       pool,
		Ctx:        ctx,
		CancelFunc: cancelFunc,
	}
}
