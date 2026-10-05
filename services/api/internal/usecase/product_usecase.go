package usecase

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"

	"pos-saas/api/internal/domain"
)

type ProductUsecase struct {
	repo domain.ProductRepository
}

func NewProductUsecase(repo domain.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo: repo}
}

func (u *ProductUsecase) Create(ctx context.Context, tenantID uuid.UUID, in domain.CreateProductInput) (*domain.Product, error) {
	if tenantID == uuid.Nil {
		return nil, domain.ErrInvalidTenantID
	}

	trimmedName := strings.TrimSpace(in.Name)
	if trimmedName == "" {
		return nil, domain.ErrInvalidName
	}

	if in.Price < 0 {
		return nil, domain.ErrInvalidPrice
	}

	if in.COGS < 0 {
		return nil, domain.ErrInvalidCOGS
	}

	trackStock := true
	if in.TrackStock != nil {
		trackStock = *in.TrackStock
	}

	unit := strings.TrimSpace(in.Unit)
	if unit == "" {
		unit = "pcs"
	}

	now := time.Now().UTC()
	prod := &domain.Product{
		ID:         uuid.New(),
		TenantID:   tenantID,
		CategoryID: in.CategoryID,
		SKU:        strings.TrimSpace(in.SKU),
		Barcode:    strings.TrimSpace(in.Barcode),
		Name:       trimmedName,
		COGS:       in.COGS,
		Price:      in.Price,
		TrackStock: trackStock,
		Unit:       unit,
		IsActive:   true,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	if err := u.repo.Create(ctx, prod); err != nil {
		return nil, err
	}

	return prod, nil
}

func (u *ProductUsecase) GetByID(ctx context.Context, tenantID, id uuid.UUID) (*domain.Product, error) {
	if tenantID == uuid.Nil {
		return nil, domain.ErrInvalidTenantID
	}
	if id == uuid.Nil {
		return nil, domain.ErrInvalidProductID
	}

	return u.repo.GetByID(ctx, tenantID, id)
}

func (u *ProductUsecase) List(ctx context.Context, tenantID uuid.UUID, limit, offset int) ([]domain.Product, int64, error) {
	if tenantID == uuid.Nil {
		return nil, 0, domain.ErrInvalidTenantID
	}

	return u.repo.List(ctx, tenantID, limit, offset)
}

func (u *ProductUsecase) Update(ctx context.Context, tenantID, id uuid.UUID, in domain.UpdateProductInput) (*domain.Product, error) {
	if tenantID == uuid.Nil {
		return nil, domain.ErrInvalidTenantID
	}
	if id == uuid.Nil {
		return nil, domain.ErrInvalidProductID
	}

	existing, err := u.repo.GetByID(ctx, tenantID, id)
	if err != nil {
		return nil, err
	}

	if in.Name != nil {
		trimmed := strings.TrimSpace(*in.Name)
		if trimmed == "" {
			return nil, domain.ErrInvalidName
		}
		existing.Name = trimmed
	}

	if in.Price != nil {
		if *in.Price < 0 {
			return nil, domain.ErrInvalidPrice
		}
		existing.Price = *in.Price
	}

	if in.COGS != nil {
		if *in.COGS < 0 {
			return nil, domain.ErrInvalidCOGS
		}
		existing.COGS = *in.COGS
	}

	if in.CategoryID != nil {
		existing.CategoryID = in.CategoryID
	}

	if in.SKU != nil {
		existing.SKU = strings.TrimSpace(*in.SKU)
	}

	if in.Barcode != nil {
		existing.Barcode = strings.TrimSpace(*in.Barcode)
	}

	if in.TrackStock != nil {
		existing.TrackStock = *in.TrackStock
	}

	if in.Unit != nil {
		trimmedUnit := strings.TrimSpace(*in.Unit)
		if trimmedUnit != "" {
			existing.Unit = trimmedUnit
		}
	}

	if in.IsActive != nil {
		existing.IsActive = *in.IsActive
	}

	existing.UpdatedAt = time.Now().UTC()

	if err := u.repo.Update(ctx, existing); err != nil {
		return nil, err
	}

	return existing, nil
}

func (u *ProductUsecase) Delete(ctx context.Context, tenantID, id uuid.UUID) error {
	if tenantID == uuid.Nil {
		return domain.ErrInvalidTenantID
	}
	if id == uuid.Nil {
		return domain.ErrInvalidProductID
	}

	return u.repo.Delete(ctx, tenantID, id)
}
