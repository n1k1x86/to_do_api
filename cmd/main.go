package main

import (
	"fmt"
	"log"
	"todoapi/internal/config"
	"todoapi/internal/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error while loading config: %v", err)
	}
	if err = database.UpMigrations(&cfg.Database); err != nil {
		log.Fatalf("error while setting migrations: %v", err)
	}
	fmt.Println("lets start")
}
