package usecase_test

import (
	"context"
	"testing"

	"github.com/google/uuid"

	"pos-saas/api/internal/domain"
	"pos-saas/api/internal/usecase"
)

type mockProductRepo struct {
	products map[string]*domain.Product
}

func newMockProductRepo() *mockProductRepo {
	return &mockProductRepo{
		products: make(map[string]*domain.Product),
	}
}

func (m *mockProductRepo) key(tenantID, id uuid.UUID) string {
	return tenantID.String() + ":" + id.String()
}

func (m *mockProductRepo) Create(ctx context.Context, p *domain.Product) error {
	m.products[m.key(p.TenantID, p.ID)] = p
	return nil
}

func (m *mockProductRepo) GetByID(ctx context.Context, tenantID, id uuid.UUID) (*domain.Product, error) {
	p, ok := m.products[m.key(tenantID, id)]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return p, nil
}

func (m *mockProductRepo) GetBySKU(ctx context.Context, tenantID uuid.UUID, sku string) (*domain.Product, error) {
	for _, p := range m.products {
		if p.TenantID == tenantID && p.SKU == sku {
			return p, nil
		}
	}
	return nil, domain.ErrNotFound
}

func (m *mockProductRepo) GetByBarcode(ctx context.Context, tenantID uuid.UUID, barcode string) (*domain.Product, error) {
	for _, p := range m.products {
		if p.TenantID == tenantID && p.Barcode == barcode {
			return p, nil
		}
	}
	return nil, domain.ErrNotFound
}

func (m *mockProductRepo) List(ctx context.Context, tenantID uuid.UUID, limit, offset int) ([]domain.Product, int64, error) {
	var list []domain.Product
	for _, p := range m.products {
		if p.TenantID == tenantID {
			list = append(list, *p)
		}
	}
	return list, int64(len(list)), nil
}

func (m *mockProductRepo) Update(ctx context.Context, p *domain.Product) error {
	k := m.key(p.TenantID, p.ID)
	if _, ok := m.products[k]; !ok {
		return domain.ErrNotFound
	}
	m.products[k] = p
	return nil
}

func (m *mockProductRepo) Delete(ctx context.Context, tenantID, id uuid.UUID) error {
	k := m.key(tenantID, id)
	if _, ok := m.products[k]; !ok {
		return domain.ErrNotFound
	}
	delete(m.products, k)
	return nil
}

func TestProductUsecase_Create(t *testing.T) {
	repo := newMockProductRepo()
	uc := usecase.NewProductUsecase(repo)
	tenantID := uuid.New()

	t.Run("success create and calculate margin", func(t *testing.T) {
		prod, err := uc.Create(context.Background(), tenantID, domain.CreateProductInput{
			Name:  "Kopi Susu Gula Aren",
			SKU:   "BEV-001",
			COGS:  8000,
			Price: 18000,
		})

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if prod.Name != "Kopi Susu Gula Aren" {
			t.Errorf("expected name Kopi Susu Gula Aren, got %s", prod.Name)
		}
		if prod.Margin() != 10000 {
			t.Errorf("expected margin 10000, got %d", prod.Margin())
		}
		expectedPct := (10000.0 / 18000.0) * 100
		if prod.MarginPercentage() != expectedPct {
			t.Errorf("expected margin pct %f, got %f", expectedPct, prod.MarginPercentage())
		}
		if prod.Unit != "pcs" {
			t.Errorf("expected default unit pcs, got %s", prod.Unit)
		}
		if !prod.TrackStock {
			t.Errorf("expected default track_stock true")
		}
	})

	t.Run("fail when name is empty", func(t *testing.T) {
		_, err := uc.Create(context.Background(), tenantID, domain.CreateProductInput{
			Name:  "   ",
			Price: 10000,
		})
		if err != domain.ErrInvalidName {
			t.Fatalf("expected ErrInvalidName, got %v", err)
		}
	})

	t.Run("fail when price is negative", func(t *testing.T) {
		_, err := uc.Create(context.Background(), tenantID, domain.CreateProductInput{
			Name:  "Product Neg Price",
			Price: -100,
		})
		if err != domain.ErrInvalidPrice {
			t.Fatalf("expected ErrInvalidPrice, got %v", err)
		}
	})

	t.Run("fail when COGS is negative", func(t *testing.T) {
		_, err := uc.Create(context.Background(), tenantID, domain.CreateProductInput{
			Name:  "Product Neg COGS",
			Price: 1000,
			COGS:  -50,
		})
		if err != domain.ErrInvalidCOGS {
			t.Fatalf("expected ErrInvalidCOGS, got %v", err)
		}
	})
}

func TestProductUsecase_Update(t *testing.T) {
	repo := newMockProductRepo()
	uc := usecase.NewProductUsecase(repo)
	tenantID := uuid.New()

	prod, _ := uc.Create(context.Background(), tenantID, domain.CreateProductInput{
		Name:  "Original",
		Price: 10000,
		COGS:  5000,
	})

	newName := "Updated Name"
	newPrice := int64(15000)

	updated, err := uc.Update(context.Background(), tenantID, prod.ID, domain.UpdateProductInput{
		Name:  &newName,
		Price: &newPrice,
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if updated.Name != "Updated Name" {
		t.Errorf("expected updated name, got %s", updated.Name)
	}
	if updated.Price != 15000 {
		t.Errorf("expected updated price, got %d", updated.Price)
	}
	if updated.Margin() != 10000 {
		t.Errorf("expected new margin 10000, got %d", updated.Margin())
	}
}

func TestProductUsecase_Delete(t *testing.T) {
	repo := newMockProductRepo()
	uc := usecase.NewProductUsecase(repo)
	tenantID := uuid.New()

	prod, _ := uc.Create(context.Background(), tenantID, domain.CreateProductInput{
		Name:  "To Delete",
		Price: 5000,
	})

	err := uc.Delete(context.Background(), tenantID, prod.ID)
	if err != nil {
		t.Fatalf("expected delete success, got %v", err)
	}

	_, err = uc.GetByID(context.Background(), tenantID, prod.ID)
	if err != domain.ErrNotFound {
		t.Fatalf("expected ErrNotFound after delete, got %v", err)
	}
}
