package routes

import (
	"wallet-service/internal/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, ctrl *controllers.WalletController) {
	w := app.Group("/api/wallet")
	w.Post("/redeem", ctrl.Redeem)
	w.Get("/balance/:phone", ctrl.GetBalance)
	w.Get("/redeemed", ctrl.GetRedeemed)
}
