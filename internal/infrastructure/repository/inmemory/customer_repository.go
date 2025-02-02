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
	Customers map[id.ID]*customer.Customer
}

func NewCustomerRepository() *customerRepository {
	return &customerRepository{
		Customers: make(map[id.ID]*customer.Customer),
	}
}

func (r *customerRepository) Save(ctx context.Context, customer *customer.Customer) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	r.Customers[customer.ID] = customer
	return nil
}

func (r *customerRepository) FindByID(ctx context.Context, id id.ID) (*customer.Customer, *apperr.AppErr) {
	r.RLock()
	defer r.RUnlock()

	if _, ok := r.Customers[id]; !ok {
		return nil, apperr.New("not found")
	}

	return r.Customers[id], nil
}
