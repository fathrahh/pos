package http

import (
	"net/http"

	"pos-saas/api/pkg/response"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]any{
		"status":  "healthy",
		"service": "pos-saas-api",
	})
}
