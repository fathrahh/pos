package middleware

import (
	"context"
	"net/http"

	"pos-saas/api/pkg/response"
)

type contextKey string

const (
	TenantIDKey contextKey = "tenant_id"
	OutletIDKey contextKey = "outlet_id"
	UserIDKey   contextKey = "user_id"
)

func RequireTenant(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tenantID := r.Header.Get("X-Tenant-ID")
		if tenantID == "" {
			response.Error(w, http.StatusBadRequest, "Missing X-Tenant-ID header")
			return
		}

		ctx := context.WithValue(r.Context(), TenantIDKey, tenantID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetTenantID(ctx context.Context) string {
	if val, ok := ctx.Value(TenantIDKey).(string); ok {
		return val
	}
	return ""
}
