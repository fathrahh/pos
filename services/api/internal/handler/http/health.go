package http

import (
	"github.com/gofiber/fiber/v2"

	"pos-saas/api/pkg/response"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Check(c *fiber.Ctx) error {
	return response.JSON(c, fiber.StatusOK, map[string]any{
		"status":  "healthy",
		"service": "pos-saas-api",
	})
}
