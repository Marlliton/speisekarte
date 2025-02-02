package inmemory

import (
	"context"
	"sync"

	"github.com/Marlliton/speisekarte/internal/entity/customer"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/Marlliton/speisekarte/pkg/id"
)

type customerRespository struct {
	sync.RWMutex
	customers map[id.ID]*customer.Customer
}

func NewCustomerRepository() *customerRespository {
	return &customerRespository{
		customers: make(map[id.ID]*customer.Customer),
	}
}

func (r *customerRespository) Save(ctx context.Context, customer *customer.Customer) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	r.customers[customer.ID] = customer
	return nil
}

func (r *customerRespository) FindByID(ctx context.Context, id id.ID) (*customer.Customer, *apperr.AppErr) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.customers[id]; !ok {
		return nil, apperr.New("not found")
	}

	return r.customers[id], nil
}
