package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todoapi/internal/config"
	"todoapi/internal/database"
	"todoapi/internal/server"
	tasksStorager "todoapi/internal/services/tasks/storager"
	usersStorager "todoapi/internal/services/users/storager"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error while loading config: %v", err)
	}
	if err = database.UpMigrations(&cfg.Database); err != nil {
		log.Fatalf("error while setting migrations: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool, err := database.NewPool(ctx, &cfg.Database)
	if err != nil {
		log.Fatalf("error while creating pgx connection pool: %v", err)
	}
	defer pool.Close()

	usersRepo := usersStorager.NewUsersRepo(pool)
	tasksRepo := tasksStorager.NewTasksRepo(pool)
	server := server.NewHTTPServer(ctx, &cfg.Server, usersRepo, tasksRepo)

	go func() {
		defer func() {
			r := recover()
			if r != nil {
				log.Printf("panic was recovered while running server: %s", r)
			}
		}()
		err = server.Run()
		if err != nil {
			log.Printf("error while running server: %v", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGINT)
	<-sig

	graceCtx, graceCancel := context.WithTimeout(ctx, time.Second*5)
	defer graceCancel()

	server.Shutdown(graceCtx)
	log.Println("app was closed successfully")
}
