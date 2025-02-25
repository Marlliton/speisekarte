package repository

import (
	"context"

	"github.com/Marlliton/speisekarte/internal/entity/product"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/Marlliton/speisekarte/pkg/id"
)

type ProductRepository interface {
	Create(ctx context.Context, product *product.Product) *apperr.AppErr
	FindByID(ctx context.Context, id id.ID) (*product.Product, *apperr.AppErr)
	FindAll(ctx context.Context) ([]*product.Product, *apperr.AppErr)
	Delete(ctx context.Context, id id.ID) *apperr.AppErr
	Update(ctx context.Context, id id.ID, product *product.Product) *apperr.AppErr
}
