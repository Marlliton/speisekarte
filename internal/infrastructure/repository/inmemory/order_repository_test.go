package inmemory

import (
	"context"
	"testing"

	"github.com/Marlliton/speisekarte/internal/entity/order"
	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestOrderRepository(t *testing.T) {
	var repo *inMemoryOrderRepository
	ctx := context.Background()

	setup := func() {
		repo = NewInMemoryOrderRepository()
	}

	t.Run("should create an order", func(t *testing.T) {
		setup()
		order := &order.Order{
			ID:         id.New(),
			CustomerID: id.New(),
		}

		err := repo.Create(ctx, order)
		assert.Nil(t, err)

		foundOrders, err := repo.FindAllByCustomerID(order.CustomerID)
		assert.Nil(t, err)
		assert.Len(t, foundOrders, 1)
		assert.Equal(t, order.ID, foundOrders[0].ID)
		assert.Equal(t, order.CustomerID, foundOrders[0].CustomerID)
	})

	t.Run("should find all orders by customer ID", func(t *testing.T) {
		setup()
		customerID := id.New()
		order1 := &order.Order{
			ID:         id.New(),
			CustomerID: customerID,
		}
		order2 := &order.Order{
			ID:         id.New(),
			CustomerID: customerID,
		}
		order3 := &order.Order{
			ID:         id.New(),
			CustomerID: id.New(), // Different customer
		}

		repo.Create(ctx, order1)
		repo.Create(ctx, order2)
		repo.Create(ctx, order3)

		foundOrders, err := repo.FindAllByCustomerID(customerID)
		assert.Nil(t, err)
		assert.Len(t, foundOrders, 2)
		assert.Equal(t, order1.ID, foundOrders[0].ID)
		assert.Equal(t, order2.ID, foundOrders[1].ID)
	})

	t.Run("should return empty list if no orders found for customer ID", func(t *testing.T) {
		setup()
		customerID := id.New()

		foundOrders, err := repo.FindAllByCustomerID(customerID)
		assert.Nil(t, err)
		assert.Empty(t, foundOrders)
	})
}
