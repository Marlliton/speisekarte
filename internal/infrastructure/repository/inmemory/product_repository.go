package inmemory

import (
	"context"
	"fmt"
	"sync"

	"github.com/Marlliton/speisekarte/internal/entity/product"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/Marlliton/speisekarte/pkg/id"
)

type inMemoryProductRepository struct {
	sync.RWMutex
	products map[id.ID]*product.Product
}

func NewInMemoryProductRepository() *inMemoryProductRepository {
	return &inMemoryProductRepository{products: map[id.ID]*product.Product{}}
}

func (r *inMemoryProductRepository) Create(ctx context.Context, product *product.Product) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	r.products[product.ID] = product
	return nil
}

func (r *inMemoryProductRepository) FindByID(ctx context.Context, id id.ID) (*product.Product, *apperr.AppErr) {
	r.RLock()
	defer r.RUnlock()

	if _, exists := r.products[id]; !exists {
		return nil, apperr.New("not found").WithCode(apperr.NOT_FOUND)
	}

	return r.products[id], nil
}

func (r *inMemoryProductRepository) FindAll(ctx context.Context) ([]*product.Product, *apperr.AppErr) {
	r.RLock()
	defer r.RUnlock()

	products := make([]*product.Product, 0, len(r.products))
	for _, c := range r.products {
		products = append(products, c)
	}

	return products, nil
}

func (r *inMemoryProductRepository) Delete(ctx context.Context, id id.ID) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.products[id]; !exists {
		return apperr.New("not found")
	}

	delete(r.products, id)
	return nil
}

func (r *inMemoryProductRepository) Update(ctx context.Context, id id.ID, product *product.Product) *apperr.AppErr {
	if product == nil {
		return apperr.New("product cannot be nil")
	}

	r.Lock()
	defer r.Unlock()

	_, exists := r.products[id]
	if !exists {
		return apperr.New(fmt.Sprintf("product with ID %s not found", id))
	}

	r.products[id] = product

	return nil
}
