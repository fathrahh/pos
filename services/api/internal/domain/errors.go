package domain

import "errors"

var (
	ErrNotFound         = errors.New("product not found")
	ErrDuplicateSKU     = errors.New("product with this sku already exists")
	ErrDuplicateBarcode = errors.New("product with this barcode already exists")
	ErrInvalidName      = errors.New("product name is required")
	ErrInvalidPrice     = errors.New("price cannot be negative")
	ErrInvalidCOGS      = errors.New("cogs cannot be negative")
	ErrInvalidTenantID  = errors.New("invalid or missing tenant id")
	ErrInvalidProductID = errors.New("invalid product id")
)
