package http

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"pos-saas/api/internal/domain"
	"pos-saas/api/internal/middleware"
	"pos-saas/api/pkg/response"
)

type ProductHandler struct {
	usecase domain.ProductUsecase
}

func NewProductHandler(usecase domain.ProductUsecase) *ProductHandler {
	return &ProductHandler{usecase: usecase}
}

func (h *ProductHandler) getTenantUUID(c *fiber.Ctx) (uuid.UUID, error) {
	tenantIDStr := middleware.GetTenantID(c.UserContext())
	if tenantIDStr == "" {
		tenantIDStr = c.Get("X-Tenant-ID")
	}
	return uuid.Parse(tenantIDStr)
}

func (h *ProductHandler) Create(c *fiber.Ctx) error {
	tenantID, err := h.getTenantUUID(c)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid or missing X-Tenant-ID header")
	}

	var in domain.CreateProductInput
	if err := c.BodyParser(&in); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	prod, err := h.usecase.Create(c.UserContext(), tenantID, in)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrInvalidName),
			errors.Is(err, domain.ErrInvalidPrice),
			errors.Is(err, domain.ErrInvalidCOGS):
			return response.Error(c, fiber.StatusBadRequest, err.Error())
		case errors.Is(err, domain.ErrDuplicateSKU),
			errors.Is(err, domain.ErrDuplicateBarcode):
			return response.Error(c, fiber.StatusConflict, err.Error())
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to create product: "+err.Error())
		}
	}

	return response.JSON(c, fiber.StatusCreated, prod)
}

func (h *ProductHandler) List(c *fiber.Ctx) error {
	tenantID, err := h.getTenantUUID(c)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid or missing X-Tenant-ID header")
	}

	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	items, total, err := h.usecase.List(c.UserContext(), tenantID, limit, offset)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list products: "+err.Error())
	}

	return response.JSON(c, fiber.StatusOK, fiber.Map{
		"items":  items,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	})
}

func (h *ProductHandler) GetByID(c *fiber.Ctx) error {
	tenantID, err := h.getTenantUUID(c)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid or missing X-Tenant-ID header")
	}

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid product id format")
	}

	prod, err := h.usecase.GetByID(c.UserContext(), tenantID, id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return response.Error(c, fiber.StatusNotFound, "Product not found")
		}
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch product: "+err.Error())
	}

	return response.JSON(c, fiber.StatusOK, prod)
}

func (h *ProductHandler) Update(c *fiber.Ctx) error {
	tenantID, err := h.getTenantUUID(c)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid or missing X-Tenant-ID header")
	}

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid product id format")
	}

	var in domain.UpdateProductInput
	if err := c.BodyParser(&in); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body: "+err.Error())
	}

	prod, err := h.usecase.Update(c.UserContext(), tenantID, id, in)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			return response.Error(c, fiber.StatusNotFound, "Product not found")
		case errors.Is(err, domain.ErrInvalidName),
			errors.Is(err, domain.ErrInvalidPrice),
			errors.Is(err, domain.ErrInvalidCOGS):
			return response.Error(c, fiber.StatusBadRequest, err.Error())
		case errors.Is(err, domain.ErrDuplicateSKU),
			errors.Is(err, domain.ErrDuplicateBarcode):
			return response.Error(c, fiber.StatusConflict, err.Error())
		default:
			return response.Error(c, fiber.StatusInternalServerError, "Failed to update product: "+err.Error())
		}
	}

	return response.JSON(c, fiber.StatusOK, prod)
}

func (h *ProductHandler) Delete(c *fiber.Ctx) error {
	tenantID, err := h.getTenantUUID(c)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid or missing X-Tenant-ID header")
	}

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid product id format")
	}

	if err := h.usecase.Delete(c.UserContext(), tenantID, id); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return response.Error(c, fiber.StatusNotFound, "Product not found")
		}
		return response.Error(c, fiber.StatusInternalServerError, "Failed to delete product: "+err.Error())
	}

	return response.JSON(c, fiber.StatusOK, fiber.Map{
		"message": "Product deleted successfully",
	})
}
