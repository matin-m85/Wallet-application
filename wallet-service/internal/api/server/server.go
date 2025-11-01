package api

import (
	"wallet-service/internal/api/controllers"
	"wallet-service/internal/api/routes"
	"wallet-service/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

func NewApp(db *gorm.DB) *fiber.App {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())

	// create service from db
	wsvc := service.NewWalletService(db)
	ctrl := controllers.NewWalletController(wsvc)
	routes.SetupRoutes(app, ctrl)
	return app
}
