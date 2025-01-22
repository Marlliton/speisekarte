package order

import (
	"testing"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestNewOrder_Success(t *testing.T) {
	customerID := id.New()
	items := []*OrderItem{
		{
			ID:        id.New(),
			ProductID: id.New(),
			Quantity:  2,
			Price:     1000,
		},
	}

	t.Run("should create an order", func(t *testing.T) {
		order, errs := New(customerID, items...)

		assert.Nil(t, errs)
		assert.NotNil(t, order)
		assert.Equal(t, customerID, order.CustomerID)
		assert.Equal(t, Pending, order.Status)
		assert.Len(t, order.Items, 1)
	})

}

func TestNewOrder_FailToCreate(t *testing.T) {
	t.Run("should fail to create whitout id", func(t *testing.T) {
		order, errs := New(id.ID{}, &OrderItem{
			ID:        id.New(),
			ProductID: id.New(),
			Quantity:  1,
			Price:     1000,
		})

		assert.Nil(t, order)
		assert.NotNil(t, errs)
	})

	t.Run("should to create whitout order item", func(t *testing.T) {
		order, errs := New(id.New())

		assert.Nil(t, order)
		assert.NotNil(t, errs)
	})
}
