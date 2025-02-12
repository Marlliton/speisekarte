package inmemory

import (
	"context"
	"sync"

	"github.com/Marlliton/speisekarte/internal/entity/order"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/Marlliton/speisekarte/pkg/id"
)

type inMemoryOrderRepository struct {
	sync.RWMutex
	orders map[id.ID]*order.Order
}

func NewInMemoryOrderRepository() *inMemoryOrderRepository {
	return &inMemoryOrderRepository{
		orders: make(map[id.ID]*order.Order, 0),
	}
}

func (r *inMemoryOrderRepository) Create(ctx context.Context, order *order.Order) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	r.orders[order.ID] = order
	return nil
}

func (r *inMemoryOrderRepository) FindAllByCustomerID(customerID id.ID) ([]*order.Order, *apperr.AppErr) {
	r.RLock()
	defer r.RUnlock()

	orders := make([]*order.Order, 0)
	for _, order := range r.orders {
		if order.CustomerID == customerID {
			orders = append(orders, order)
		}
	}

	return orders, nil
}
