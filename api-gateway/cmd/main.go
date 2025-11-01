package main

import (
	"api-gateway/config"
	api "api-gateway/internal/api/server"
	"api-gateway/internal/service"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	svc := service.NewGatewayService(cfg)
	app := api.NewApp(cfg, svc)

	log.Printf("🚀 API Gateway running on port %s\n", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("gateway stopped: %v", err)
	}
}
