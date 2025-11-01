package controllers

import (
	"api-gateway/internal/service"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type GatewayController struct {
	svc          *service.GatewayService
	walletBase   string
	discountBase string
}

func NewGatewayController(svc *service.GatewayService, walletBase, discountBase string) *GatewayController {
	return &GatewayController{
		svc:          svc,
		walletBase:   walletBase,
		discountBase: discountBase,
	}
}

func (g *GatewayController) Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok", "service": "gateway"})
}

func copyHeaders(c *fiber.Ctx) http.Header {
	h := http.Header{}
	c.Request().Header.VisitAll(func(k, v []byte) {
		h.Set(string(k), string(v))
	})
	return h
}

func (g *GatewayController) forward(c *fiber.Ctx, target string) error {
	origPath := c.OriginalURL() // includes path and query
	// context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	headers := copyHeaders(c)
	method := c.Method()
	body := c.Body()

	var base string
	if target == "wallet" {
		base = g.walletBase
	} else {
		base = g.discountBase
	}
	status, respHeaders, respBody, err := g.svc.DoForward(ctx, method, base, origPath, headers, body)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
	}
	// copy back a couple headers
	if ct := respHeaders.Get("Content-Type"); ct != "" {
		c.Set("Content-Type", ct)
	}
	return c.Status(status).Send(respBody)
}

func (g *GatewayController) WalletProxy(c *fiber.Ctx) error {
	if !strings.HasPrefix(c.OriginalURL(), "/api/wallet") && !strings.HasPrefix(c.OriginalURL(), "/api/v1/wallets") {
		return c.Status(400).JSON(fiber.Map{"error": "invalid wallet path"})
	}
	return g.forward(c, "wallet")
}

func (g *GatewayController) DiscountProxy(c *fiber.Ctx) error {
	if !strings.HasPrefix(c.OriginalURL(), "/api/discount") && !strings.HasPrefix(c.OriginalURL(), "/api/v1/discounts") {
		return c.Status(400).JSON(fiber.Map{"error": "invalid discount path"})
	}
	return g.forward(c, "discount")
}
