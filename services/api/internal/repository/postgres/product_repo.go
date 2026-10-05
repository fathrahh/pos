package postgres

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"pos-saas/api/internal/domain"
)

type ProductModel struct {
	ID         uuid.UUID  `gorm:"type:uuid;primaryKey"`
	TenantID   uuid.UUID  `gorm:"type:uuid;not null;index"`
	CategoryID *uuid.UUID `gorm:"type:uuid;index"`
	SKU        string     `gorm:"type:varchar(50)"`
	Barcode    string     `gorm:"type:varchar(100)"`
	Name       string     `gorm:"type:varchar(255);not null"`
	COGS       int64      `gorm:"type:bigint;not null;default:0"`
	Price      int64      `gorm:"type:bigint;not null;default:0"`
	TrackStock bool       `gorm:"type:boolean;not null;default:true"`
	Unit       string     `gorm:"type:varchar(20);not null;default:'pcs'"`
	IsActive   bool       `gorm:"type:boolean;not null;default:true"`
	CreatedAt  time.Time  `gorm:"not null"`
	UpdatedAt  time.Time  `gorm:"not null"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (ProductModel) TableName() string {
	return "products"
}

func toDomainProduct(m *ProductModel) *domain.Product {
	var deletedAt *time.Time
	if m.DeletedAt.Valid {
		deletedAt = &m.DeletedAt.Time
	}
	return &domain.Product{
		ID:         m.ID,
		TenantID:   m.TenantID,
		CategoryID: m.CategoryID,
		SKU:        m.SKU,
		Barcode:    m.Barcode,
		Name:       m.Name,
		COGS:       m.COGS,
		Price:      m.Price,
		TrackStock: m.TrackStock,
		Unit:       m.Unit,
		IsActive:   m.IsActive,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
		DeletedAt:  deletedAt,
	}
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(ctx context.Context, p *domain.Product) error {
	model := ProductModel{
		ID:         p.ID,
		TenantID:   p.TenantID,
		CategoryID: p.CategoryID,
		SKU:        p.SKU,
		Barcode:    p.Barcode,
		Name:       p.Name,
		COGS:       p.COGS,
		Price:      p.Price,
		TrackStock: p.TrackStock,
		Unit:       p.Unit,
		IsActive:   p.IsActive,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}

	err := r.db.WithContext(ctx).Create(&model).Error
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "idx_products_tenant_sku") ||
			(strings.Contains(err.Error(), "duplicate key") && strings.Contains(err.Error(), "sku")) {
			return domain.ErrDuplicateSKU
		}
		if strings.Contains(strings.ToLower(err.Error()), "idx_products_tenant_barcode") ||
			(strings.Contains(err.Error(), "duplicate key") && strings.Contains(err.Error(), "barcode")) {
			return domain.ErrDuplicateBarcode
		}
		return err
	}
	return nil
}

func (r *ProductRepository) GetByID(ctx context.Context, tenantID, id uuid.UUID) (*domain.Product, error) {
	var model ProductModel
	err := r.db.WithContext(ctx).
		Where("tenant_id = ? AND id = ?", tenantID, id).
		First(&model).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return toDomainProduct(&model), nil
}

func (r *ProductRepository) GetBySKU(ctx context.Context, tenantID uuid.UUID, sku string) (*domain.Product, error) {
	if sku == "" {
		return nil, domain.ErrNotFound
	}
	var model ProductModel
	err := r.db.WithContext(ctx).
		Where("tenant_id = ? AND sku = ?", tenantID, sku).
		First(&model).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return toDomainProduct(&model), nil
}

func (r *ProductRepository) GetByBarcode(ctx context.Context, tenantID uuid.UUID, barcode string) (*domain.Product, error) {
	if barcode == "" {
		return nil, domain.ErrNotFound
	}
	var model ProductModel
	err := r.db.WithContext(ctx).
		Where("tenant_id = ? AND barcode = ?", tenantID, barcode).
		First(&model).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return toDomainProduct(&model), nil
}

func (r *ProductRepository) List(ctx context.Context, tenantID uuid.UUID, limit, offset int) ([]domain.Product, int64, error) {
	var models []ProductModel
	var total int64

	query := r.db.WithContext(ctx).Model(&ProductModel{}).Where("tenant_id = ?", tenantID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}

	err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&models).Error
	if err != nil {
		return nil, 0, err
	}

	products := make([]domain.Product, len(models))
	for i, m := range models {
		products[i] = *toDomainProduct(&m)
	}

	return products, total, nil
}

func (r *ProductRepository) Update(ctx context.Context, p *domain.Product) error {
	updates := map[string]any{
		"category_id": p.CategoryID,
		"sku":         p.SKU,
		"barcode":     p.Barcode,
		"name":        p.Name,
		"cogs":        p.COGS,
		"price":       p.Price,
		"track_stock": p.TrackStock,
		"unit":        p.Unit,
		"is_active":   p.IsActive,
		"updated_at":  time.Now().UTC(),
	}

	res := r.db.WithContext(ctx).
		Model(&ProductModel{}).
		Where("tenant_id = ? AND id = ?", p.TenantID, p.ID).
		Updates(updates)

	if res.Error != nil {
		if strings.Contains(strings.ToLower(res.Error.Error()), "idx_products_tenant_sku") {
			return domain.ErrDuplicateSKU
		}
		if strings.Contains(strings.ToLower(res.Error.Error()), "idx_products_tenant_barcode") {
			return domain.ErrDuplicateBarcode
		}
		return res.Error
	}
	if res.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (r *ProductRepository) Delete(ctx context.Context, tenantID, id uuid.UUID) error {
	res := r.db.WithContext(ctx).
		Where("tenant_id = ? AND id = ?", tenantID, id).
		Delete(&ProductModel{})

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
}
