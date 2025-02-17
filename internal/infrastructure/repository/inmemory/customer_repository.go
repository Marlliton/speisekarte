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

func (r *inMemoryCustomerRepository) FindByPhone(ctx context.Context, phone string) (*customer.Customer, *apperr.AppErr) {
	r.RLock()
	defer r.RUnlock()

	for _, c := range r.customers {
		if c.Phone == phone {
			return c, nil
		}
	}

	return nil, apperr.New("not found")
}
