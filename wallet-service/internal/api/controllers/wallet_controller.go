package controllers

import (
	"net/http"

	"wallet-service/internal/service"

	"github.com/gofiber/fiber/v2"
)

type WalletController struct {
	svc *service.WalletService
}

func NewWalletController(svc *service.WalletService) *WalletController {
	return &WalletController{svc: svc}
}

type redeemReq struct {
	Phone     string `json:"phone"`
	Amount    int64  `json:"amount"`
	Reference string `json:"reference"`
}

func (c *WalletController) Redeem(ctx *fiber.Ctx) error {
	var req redeemReq
	if err := ctx.BodyParser(&req); err != nil || req.Phone == "" || req.Amount == 0 {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}
	txn, err := c.svc.ApplyTopUp(ctx.Context(), req.Phone, req.Amount, req.Reference)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(txn)
}

func (c *WalletController) GetBalance(ctx *fiber.Ctx) error {
	phone := ctx.Params("phone")
	if phone == "" {
		return ctx.Status(400).JSON(fiber.Map{"error": "phone required"})
	}
	b, err := c.svc.GetBalance(ctx.Context(), phone)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{"phone": phone, "balance": b})
}

func (c *WalletController) GetRedeemed(ctx *fiber.Ctx) error {
	list, err := c.svc.ListAllWallets()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(list)
}
