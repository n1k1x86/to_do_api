package main

import (
	"log"
	"to_do/config"
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
}
