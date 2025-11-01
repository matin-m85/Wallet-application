package api

import (
	"api-gateway/config"
	"api-gateway/internal/api/controllers"
	"api-gateway/internal/api/routes"
	"api-gateway/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewApp(cfg *config.Config, svc *service.GatewayService) *fiber.App {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())

	ctrl := controllers.NewGatewayController(svc, cfg.WalletURL, cfg.DiscountURL)
	routes.RegisterRoutes(app, ctrl)
	return app
}
