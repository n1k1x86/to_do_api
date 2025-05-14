package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"to_do/config"
	"to_do/internal/services/app"
	"to_do/internal/services/database"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = database.RunMigrations(cfg.ToDoDBConfig)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	pool, err := database.CreatePool(ctx, cfg.ToDoDBConfig)
	if err != nil {
		log.Fatal(err)
	}
	app := app.CreateNewApp(ctx, cancel, pool)
	go app.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	log.Println("Shutting down application...")
	app.Close()
}
