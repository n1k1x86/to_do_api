package main

import (
	"context"
	"log"
	"todoapi/internal/config"
	"todoapi/internal/database"
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

}
