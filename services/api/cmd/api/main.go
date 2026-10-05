package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"pos-saas/api/internal/config"
	handler "pos-saas/api/internal/handler/http"
	"pos-saas/api/internal/middleware"
	"pos-saas/api/pkg/response"
)

func main() {
	cfg := config.Load()

	app := fiber.New(fiber.Config{
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:*,https://*",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Tenant-ID, X-Outlet-ID",
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	healthHandler := handler.NewHealthHandler()
	app.Get("/health", healthHandler.Check)

	// API v1
	v1 := app.Group("/api/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return response.JSON(c, fiber.StatusOK, map[string]string{
			"message": "POS SaaS API v1",
		})
	})

	// Protected tenant-scoped group
	protected := v1.Group("/", middleware.RequireTenant())
	protected.Get("/me", func(c *fiber.Ctx) error {
		tenantID := middleware.GetTenantID(c.UserContext())
		return response.JSON(c, fiber.StatusOK, map[string]string{
			"tenant_id": tenantID,
		})
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		addr := fmt.Sprintf(":%s", cfg.Port)
		log.Printf("POS SaaS Backend (Fiber) running on port %s (%s)", addr, cfg.Environment)
		if err := app.Listen(addr); err != nil {
			log.Fatalf("Server startup failed: %v", err)
		}
	}()

	<-quit
	log.Println("Shutting down server gracefully...")

	if err := app.ShutdownWithTimeout(10 * time.Second); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited successfully")
}
