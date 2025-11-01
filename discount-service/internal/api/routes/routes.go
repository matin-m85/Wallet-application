package routes

import (
	"discount-servise/internal/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, ctrl *controllers.DiscountController) {
	v1 := app.Group("/api/discount")
	v1.Post("/create", ctrl.Create)
	v1.Get("/:code", ctrl.Get)
	v1.Get("/list", ctrl.List)
}
