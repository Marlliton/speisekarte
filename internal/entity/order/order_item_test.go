package order

import (
	"testing"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestOrderItem_New(t *testing.T) {
	orderID := id.New()
	productID := id.New()
	quantity := 2
	price := 100

	t.Run("should create an order item", func(t *testing.T) {
		orderItem, errs := OrderNewItem(orderID, productID, quantity, price)

		assert.Nil(t, errs)
		assert.NotNil(t, orderItem)
		assert.Equal(t, orderID, orderItem.OrderID)
		assert.Equal(t, productID, orderItem.ProductID)
		assert.Equal(t, quantity, orderItem.Quantity)
		assert.Equal(t, price, orderItem.Price)
	})

	t.Run("should create an order item with minimum price and quantity", func(t *testing.T) {
		orderItem, errs := OrderNewItem(orderID, productID, 1, 0)

		assert.Nil(t, errs)
		assert.NotNil(t, orderItem)
		assert.Equal(t, orderID, orderItem.OrderID)
		assert.Equal(t, productID, orderItem.ProductID)
		assert.Equal(t, 1, orderItem.Quantity)
		assert.Equal(t, 0, orderItem.Price)
	})
}

func TestOrderItem_AddOn(t *testing.T) {
	item := OrderItem{
		ID:        id.New(),
		OrderID:   id.New(),
		ProductID: id.New(),
		Quantity:  2,
		Price:     100,
	}
	addOn, _ := OrderItemNewAddOn(item.ID, "Extra Cheese", 100, 2)
	assert.NotNil(t, addOn)

	orderItem, errs := OrderNewItem(item.OrderID, item.ProductID, item.Quantity, item.Price, addOn)

	assert.Nil(t, errs)
	assert.NotNil(t, orderItem)
	assert.Len(t, orderItem.AddOns, 1)
}

func TestOrderItem_FailToCreate(t *testing.T) {
	orderID := id.New()
	productID := id.New()

	t.Run("should fail to create without OrderID", func(t *testing.T) {
		orderItem, errs := OrderNewItem(id.ID{}, productID, 2, 100)

		assert.NotNil(t, errs)
		assert.Nil(t, orderItem)
	})

	t.Run("should fail to create without ProductID", func(t *testing.T) {
		orderItem, errs := OrderNewItem(orderID, id.ID{}, 2, 100)

		assert.NotNil(t, errs)
		assert.Nil(t, orderItem)
	})

	t.Run("should fail to create with zero Quantity", func(t *testing.T) {
		orderItem, errs := OrderNewItem(orderID, productID, 0, 100)

		assert.NotNil(t, errs)
		assert.Nil(t, orderItem)
	})

	t.Run("should fail to create with negative Quantity", func(t *testing.T) {
		orderItem, errs := OrderNewItem(orderID, productID, -1, 100)

		assert.NotNil(t, errs)
		assert.Nil(t, orderItem)
	})

	t.Run("should fail to create with negative Price", func(t *testing.T) {
		orderItem, errs := OrderNewItem(orderID, productID, 2, -100)

		assert.NotNil(t, errs)
		assert.Nil(t, orderItem)
	})
}
