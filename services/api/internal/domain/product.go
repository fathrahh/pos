package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID         uuid.UUID  `json:"id"`
	TenantID   uuid.UUID  `json:"tenant_id"`
	CategoryID *uuid.UUID `json:"category_id,omitempty"`
	SKU        string     `json:"sku"`
	Barcode    string     `json:"barcode"`
	Name       string     `json:"name"`
	COGS       int64      `json:"cogs"`  // Modal/HPP in integer minor units (Rupiah)
	Price      int64      `json:"price"` // Selling price in integer minor units
	TrackStock bool       `json:"track_stock"`
	Unit       string     `json:"unit"`
	IsActive   bool       `json:"is_active"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

// Margin returns net profit margin per unit in integer minor units
func (p *Product) Margin() int64 {
	return p.Price - p.COGS
}

// MarginPercentage returns gross margin percentage (0-100)
func (p *Product) MarginPercentage() float64 {
	if p.Price <= 0 {
		return 0
	}
	return float64(p.Margin()) / float64(p.Price) * 100
}

type ProductRepository interface {
	Create(ctx context.Context, p *Product) error
	GetByID(ctx context.Context, tenantID, id uuid.UUID) (*Product, error)
	GetBySKU(ctx context.Context, tenantID uuid.UUID, sku string) (*Product, error)
	GetByBarcode(ctx context.Context, tenantID uuid.UUID, barcode string) (*Product, error)
	List(ctx context.Context, tenantID uuid.UUID, limit, offset int) ([]Product, int64, error)
	Update(ctx context.Context, p *Product) error
	Delete(ctx context.Context, tenantID, id uuid.UUID) error
}

type CreateProductInput struct {
	CategoryID *uuid.UUID `json:"category_id"`
	SKU        string     `json:"sku"`
	Barcode    string     `json:"barcode"`
	Name       string     `json:"name"`
	COGS       int64      `json:"cogs"`
	Price      int64      `json:"price"`
	TrackStock *bool      `json:"track_stock"`
	Unit       string     `json:"unit"`
}

type UpdateProductInput struct {
	CategoryID *uuid.UUID `json:"category_id"`
	SKU        *string    `json:"sku"`
	Barcode    *string    `json:"barcode"`
	Name       *string    `json:"name"`
	COGS       *int64     `json:"cogs"`
	Price      *int64     `json:"price"`
	TrackStock *bool      `json:"track_stock"`
	Unit       *string    `json:"unit"`
	IsActive   *bool      `json:"is_active"`
}

type ProductUsecase interface {
	Create(ctx context.Context, tenantID uuid.UUID, in CreateProductInput) (*Product, error)
	GetByID(ctx context.Context, tenantID, id uuid.UUID) (*Product, error)
	List(ctx context.Context, tenantID uuid.UUID, limit, offset int) ([]Product, int64, error)
	Update(ctx context.Context, tenantID, id uuid.UUID, in UpdateProductInput) (*Product, error)
	Delete(ctx context.Context, tenantID, id uuid.UUID) error
}
