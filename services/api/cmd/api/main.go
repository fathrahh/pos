package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"pos-saas/api/internal/config"
	handler "pos-saas/api/internal/handler/http"
	"pos-saas/api/internal/middleware"
	"pos-saas/api/pkg/response"
)

func main() {
	cfg := config.Load()

	r := chi.NewRouter()

	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.Timeout(30 * time.Second))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-Tenant-ID", "X-Outlet-ID"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	healthHandler := handler.NewHealthHandler()
	r.Get("/health", healthHandler.Check)

	// Protected tenant-scoped API routes
	r.Route("/api/v1", func(v1 chi.Router) {
		v1.Get("/", func(w http.ResponseWriter, r *http.Request) {
			response.JSON(w, http.StatusOK, map[string]string{
				"message": "POS SaaS API v1",
			})
		})

		v1.Group(func(protected chi.Router) {
			protected.Use(middleware.RequireTenant)
			protected.Get("/me", func(w http.ResponseWriter, r *http.Request) {
				tenantID := middleware.GetTenantID(r.Context())
				response.JSON(w, http.StatusOK, map[string]string{
					"tenant_id": tenantID,
				})
			})
		})
	})

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("POS SaaS Backend running on port :%s (%s)", cfg.Port, cfg.Environment)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server startup failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited successfully")
}
