package controllers

import (
	"discount-servise/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type DiscountController struct {
	svc *service.DiscountService
}

func NewDiscountController(svc *service.DiscountService) *DiscountController {
	return &DiscountController{svc: svc}
}

type createReq struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Percent     int    `json:"percent"`
	MaxUse      int    `json:"max_use"`
}

func (c *DiscountController) Create(ctx *fiber.Ctx) error {
	var req createReq
	if err := ctx.BodyParser(&req); err != nil || req.Code == "" || req.Percent <= 0 {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}
	d, err := c.svc.Create(ctx.Context(), req.Code, req.Description, req.Percent, req.MaxUse)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(d)
}

func (c *DiscountController) Get(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	if code == "" {
		return ctx.Status(400).JSON(fiber.Map{"error": "code required"})
	}
	d, err := c.svc.GetByCode(ctx.Context(), code)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "not found"})
	}
	return ctx.JSON(d)
}

func (c *DiscountController) List(ctx *fiber.Ctx) error {
	list, err := c.svc.ListAll()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(list)
}
