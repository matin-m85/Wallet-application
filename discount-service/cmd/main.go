package main

import (
	"discount-servise/config"
	api "discount-servise/internal/api/server"
	"discount-servise/internal/repository"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	db, err := repository.NewDB(cfg)
	if err != nil {
		log.Fatalf("failed to initialize db: %v", err)
	}

	app := api.NewApp(db)
	log.Println("🚀 Discount service running on port 8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("discount service stopped: %v", err)
	}
}
