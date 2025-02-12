package inmemory

import (
	"context"
	"sync"

	"github.com/Marlliton/speisekarte/internal/entity/customer"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/Marlliton/speisekarte/pkg/id"
)

type inMemoryCustomerRepository struct {
	sync.RWMutex
	customers map[id.ID]*customer.Customer
}

func NewInMemoryCustomerRepository() *inMemoryCustomerRepository {
	return &inMemoryCustomerRepository{
		customers: make(map[id.ID]*customer.Customer),
	}
}

func (r *inMemoryCustomerRepository) Create(ctx context.Context, customer *customer.Customer) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	r.customers[customer.ID] = customer
	return nil
}

func (r *inMemoryCustomerRepository) FindByID(ctx context.Context, id id.ID) (*customer.Customer, *apperr.AppErr) {
	r.RLock()
	defer r.RUnlock()

	if _, ok := r.customers[id]; !ok {
		return nil, apperr.New("not found")
	}

	return r.customers[id], nil
}
