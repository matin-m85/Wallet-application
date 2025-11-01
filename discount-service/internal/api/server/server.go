package api

import (
	"discount-servise/internal/api/controllers"
	"discount-servise/internal/api/routes"
	"discount-servise/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

func NewApp(db *gorm.DB) *fiber.App {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())

	svc := service.NewDiscountService(db)
	ctrl := controllers.NewDiscountController(svc)
	routes.SetupRoutes(app, ctrl)
	return app
}
