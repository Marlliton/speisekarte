package inmemory

import (
	"context"
	"sync"

	"github.com/Marlliton/speisekarte/internal/entity/customer"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/Marlliton/speisekarte/pkg/id"
)

type customerRepository struct {
	sync.RWMutex
	customers map[id.ID]*customer.Customer
}

func NewCustomerRepository() *customerRepository {
	return &customerRepository{
		customers: make(map[id.ID]*customer.Customer),
	}
}

func (r *customerRepository) Create(ctx context.Context, customer *customer.Customer) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	r.customers[customer.ID] = customer
	return nil
}

func (r *customerRepository) FindByID(ctx context.Context, id id.ID) (*customer.Customer, *apperr.AppErr) {
	r.RLock()
	defer r.RUnlock()

	if _, ok := r.customers[id]; !ok {
		return nil, apperr.New("not found")
	}

	return r.customers[id], nil
}
