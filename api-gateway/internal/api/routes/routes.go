package routes

import (
	"api-gateway/internal/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, ctrl *controllers.GatewayController) {
	app.Get("/health", ctrl.Health)

	// Wallet paths (both variants)
	app.All("/api/wallet/*", ctrl.WalletProxy)
	app.All("/api/wallet", ctrl.WalletProxy)
	app.All("/api/v1/wallets/*", ctrl.WalletProxy)
	app.All("/api/v1/wallets", ctrl.WalletProxy)

	// Discount paths
	app.All("/api/discount/*", ctrl.DiscountProxy)
	app.All("/api/discount", ctrl.DiscountProxy)
	app.All("/api/v1/discounts/*", ctrl.DiscountProxy)
	app.All("/api/v1/discounts", ctrl.DiscountProxy)
}
