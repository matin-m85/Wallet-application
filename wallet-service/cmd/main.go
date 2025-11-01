package main

import (
	"log"
	"wallet-service/config"
	api "wallet-service/internal/api/server"
	"wallet-service/internal/repository"
)

func main() {
	cfg := config.LoadConfig()

	db, err := repository.NewDB(cfg)
	if err != nil {
		log.Fatalf("failed to initialize database, got error: %v", err)
	}
	// pass DB to api server (server will create service/repo from repository package)
	app := api.NewApp(db)
	log.Println("🚀 Wallet service running on port 8081")
	if err := app.Listen(":8081"); err != nil {
		log.Fatalf("wallet service stopped: %v", err)
	}
}
