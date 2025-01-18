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

	order, errs := New(customerID, items...)
	assert.Nil(t, errs)
	assert.NotNil(t, order)
	assert.Equal(t, customerID, order.CustomerID)
	assert.Equal(t, Pending, order.Status)
	assert.Len(t, order.Items, 1)
}

func TestNewOrder_ValidationFails(t *testing.T) {
	// Test with missing customerID
	order, errs := New(id.ID{}, &OrderItem{
		ID:        id.New(),
		ProductID: id.New(),
		Quantity:  1,
		Price:     1000,
	})
	assert.Nil(t, order)
	assert.NotNil(t, errs)

	// Test with empty items
	order, errs = New(id.New())
	assert.Nil(t, order)
	assert.NotNil(t, errs)
}

func TestOrderItem_Success(t *testing.T) {
	orderID := id.New()
	productID := id.New()
	item, errs := NewOrderItem(orderID, productID, 2, 500)

	assert.Nil(t, errs)
	assert.NotNil(t, item)
	assert.Equal(t, orderID, item.OrderID)
	assert.Equal(t, productID, item.ProductID)
	assert.Equal(t, 2, item.Quantity)
	assert.Equal(t, 500, item.Price)
}

func TestOrderItem_ValidationFails(t *testing.T) {
	orderID := id.New()
	productID := id.New()

	// Test with invalid quantity
	item, errs := NewOrderItem(orderID, productID, 0, 500)
	assert.Nil(t, item)
	assert.NotNil(t, errs)

	// Test with negative price
	item, errs = NewOrderItem(orderID, productID, 2, -100)
	assert.Nil(t, item)
	assert.NotNil(t, errs)
}

func TestOrder_ValidateSuccess(t *testing.T) {
	customerID := id.New()
	items := []*OrderItem{
		{
			ID:        id.New(),
			ProductID: id.New(),
			Quantity:  1,
			Price:     1000,
		},
	}

	order := &Order{
		ID:         id.New(),
		CustomerID: customerID,
		Items:      items,
		Status:     Pending,
	}

	ok, errs := order.validate()
	assert.True(t, ok)
	assert.Nil(t, errs)
}

func TestOrder_ValidateFails(t *testing.T) {
	order := &Order{
		ID:         id.New(),
		CustomerID: id.ID{}, // Invalid CustomerID
		Items:      nil,     // No items
		Status:     Pending,
	}

	ok, errs := order.validate()
	assert.False(t, ok)
	assert.NotNil(t, errs)
	assert.Len(t, errs, 2)
}
