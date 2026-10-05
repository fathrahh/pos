package middleware

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"pos-saas/api/pkg/response"
)

type contextKey string

const (
	TenantIDKey contextKey = "tenant_id"
	OutletIDKey contextKey = "outlet_id"
	UserIDKey   contextKey = "user_id"
)

func RequireTenant() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tenantID := c.Get("X-Tenant-ID")
		if tenantID == "" {
			return response.Error(c, fiber.StatusBadRequest, "Missing X-Tenant-ID header")
		}

		c.Locals(string(TenantIDKey), tenantID)
		ctx := context.WithValue(c.UserContext(), TenantIDKey, tenantID)
		c.SetUserContext(ctx)

		return c.Next()
	}
}

func GetTenantID(ctx context.Context) string {
	if val, ok := ctx.Value(TenantIDKey).(string); ok {
		return val
	}
	return ""
}
