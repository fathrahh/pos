package http_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"pos-saas/api/internal/domain"
	handler "pos-saas/api/internal/handler/http"
	"pos-saas/api/internal/middleware"
	"pos-saas/api/internal/usecase"
)

type mockRepo struct {
	items map[string]*domain.Product
}

func (m *mockRepo) Create(ctx context.Context, p *domain.Product) error {
	m.items[p.TenantID.String()+":"+p.ID.String()] = p
	return nil
}
func (m *mockRepo) GetByID(ctx context.Context, tenantID, id uuid.UUID) (*domain.Product, error) {
	p, ok := m.items[tenantID.String()+":"+id.String()]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return p, nil
}
func (m *mockRepo) GetBySKU(ctx context.Context, tenantID uuid.UUID, sku string) (*domain.Product, error) {
	return nil, domain.ErrNotFound
}
func (m *mockRepo) GetByBarcode(ctx context.Context, tenantID uuid.UUID, barcode string) (*domain.Product, error) {
	return nil, domain.ErrNotFound
}
func (m *mockRepo) List(ctx context.Context, tenantID uuid.UUID, limit, offset int) ([]domain.Product, int64, error) {
	var list []domain.Product
	for _, p := range m.items {
		if p.TenantID == tenantID {
			list = append(list, *p)
		}
	}
	return list, int64(len(list)), nil
}
func (m *mockRepo) Update(ctx context.Context, p *domain.Product) error {
	m.items[p.TenantID.String()+":"+p.ID.String()] = p
	return nil
}
func (m *mockRepo) Delete(ctx context.Context, tenantID, id uuid.UUID) error {
	delete(m.items, tenantID.String()+":"+id.String())
	return nil
}

func setupTestApp() (*fiber.App, *mockRepo, uuid.UUID) {
	app := fiber.New()
	repo := &mockRepo{items: make(map[string]*domain.Product)}
	uc := usecase.NewProductUsecase(repo)
	h := handler.NewProductHandler(uc)

	tenantID := uuid.New()

	v1 := app.Group("/api/v1", middleware.RequireTenant())
	v1.Post("/products", h.Create)
	v1.Get("/products", h.List)
	v1.Get("/products/:id", h.GetByID)
	v1.Put("/products/:id", h.Update)
	v1.Delete("/products/:id", h.Delete)

	return app, repo, tenantID
}

func TestProductHandler_CreateAndList(t *testing.T) {
	app, _, tenantID := setupTestApp()

	payload := domain.CreateProductInput{
		Name:  "Matcha Latte",
		SKU:   "DRK-002",
		COGS:  12000,
		Price: 25000,
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Tenant-ID", tenantID.String())

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201 Created, got %d", resp.StatusCode)
	}

	// Test List
	listReq := httptest.NewRequest(http.MethodGet, "/api/v1/products", nil)
	listReq.Header.Set("X-Tenant-ID", tenantID.String())

	listResp, err := app.Test(listReq)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if listResp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", listResp.StatusCode)
	}
}

func TestProductHandler_MissingTenantHeader(t *testing.T) {
	app, _, _ := setupTestApp()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/products", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400 Bad Request, got %d", resp.StatusCode)
	}
}
