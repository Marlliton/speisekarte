package order

import (
	"testing"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestNewOrderItem_Success(t *testing.T) {
	orderID := id.New()
	productID := id.New()

	t.Run("should create an order item", func(t *testing.T) {
		item, errs := OrderNewItem(id.New(), orderID, productID, 2, 1000)

		assert.Nil(t, errs)
		assert.NotNil(t, item)
		assert.Equal(t, orderID, item.OrderID)
		assert.Equal(t, productID, item.ProductID)
		assert.Equal(t, 2, item.Quantity)
		assert.Equal(t, 1000, item.Price)
	})
}

func TestNewOrderItem_FailToCreate(t *testing.T) {
	t.Run("should fail to create without order id", func(t *testing.T) {
		item, errs := OrderNewItem(id.New(), id.ID{}, id.New(), 2, 1000)

		assert.Nil(t, item)
		assert.NotNil(t, errs)
	})

	t.Run("should fail to create without product id", func(t *testing.T) {
		item, errs := OrderNewItem(id.New(), id.New(), id.ID{}, 2, 1000)

		assert.Nil(t, item)
		assert.NotNil(t, errs)
	})

	t.Run("should fail to create with invalid quantity", func(t *testing.T) {
		item, errs := OrderNewItem(id.New(), id.New(), id.New(), 0, 1000)

		assert.Nil(t, item)
		assert.NotNil(t, errs)
	})

	t.Run("should fail to create with invalid price", func(t *testing.T) {
		item, errs := OrderNewItem(id.New(), id.New(), id.New(), 2, -100)

		assert.Nil(t, item)
		assert.NotNil(t, errs)
	})
}
