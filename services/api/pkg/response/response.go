package response

import (
	"github.com/gofiber/fiber/v2"
)

type Envelope struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func JSON(c *fiber.Ctx, status int, data any) error {
	return c.Status(status).JSON(Envelope{
		Success: status >= 200 && status < 300,
		Data:    data,
	})
}

func Error(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(Envelope{
		Success: false,
		Error:   message,
	})
}
