package order

import (
	"testing"

	"github.com/Marlliton/speisekarte/internal/entity/product"
	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/Marlliton/validator/fail"
	"github.com/stretchr/testify/assert"
)

func createOrder(t *testing.T) (*Order, []*fail.Error) {
	t.Helper()
	customerID := id.New()
	order, errs := New(customerID)
	assert.Nil(t, errs)

	return order, errs
}

func createOrderItem(t *testing.T, orderID, prodID id.ID, qtdy, price int) *OrderItem {
	t.Helper()
	oItem, errs := NewOrderItem(orderID, prodID, qtdy, price)
	assert.Nil(t, errs)
	assert.NotNil(t, oItem)

	return oItem
}

func TestOrder_New(t *testing.T) {
	t.Run("create a valid order", func(t *testing.T) {
		order, errs := createOrder(t)
		assert.Nil(t, errs)
		assert.NotNil(t, order)
		assert.Equal(t, Pending, order.Status)
	})
}

func TestOrder_AddItem(t *testing.T) {
	order, _ := createOrder(t)
	assert.NotNil(t, order)
	p, _ := product.New("eggs", "fried eggs whith coke", "http://test.com", 19.65, true, id.New())
	assert.NotNil(t, p)
	item := createOrderItem(t, order.ID, p.ID, 1, p.Price)

	t.Run("add new item to order", func(t *testing.T) {
		order.AddItem(item)

		assert.Equal(t, 1, len(order.Items))
	})

	t.Run("update quantity value", func(t *testing.T) {
		item.Quantity = 4
		order.AddItem(item)

		assert.Equal(t, 1, len(order.Items))
		assert.Equal(t, 4, order.Items[0].Quantity)
	})

	t.Run("add others itens", func(t *testing.T) {
		item2 := createOrderItem(t, order.ID, id.New(), 2, p.Price)
		item3 := createOrderItem(t, order.ID, id.New(), 3, p.Price)

		order.AddItem(item2)
		order.AddItem(item3)
		assert.Equal(t, 3, len(order.Items))
		assert.Equal(t, item.ID, order.Items[0].ID)
		assert.Equal(t, item2.ID, order.Items[1].ID)
		assert.Equal(t, item3.ID, order.Items[2].ID)
		assert.Equal(t, 2, order.Items[1].Quantity)
		assert.Equal(t, 3, order.Items[2].Quantity)
	})
}

func TestOrder_Total(t *testing.T) {
	order, _ := createOrder(t)
	assert.NotNil(t, order)
	p, _ := product.New("eggs", "fried eggs whith coke", "http://test.com", 10.0, true, id.New())
	p2, _ := product.New("koke", "coke with ice", "http://test.com", 10.0, true, id.New())
	assert.NotNil(t, p)
	item := createOrderItem(t, order.ID, p.ID, 1, p.Price)
	item2 := createOrderItem(t, order.ID, p2.ID, 3, p.Price)
	order.AddItem(item)

	t.Run("get the total price order", func(t *testing.T) {
		assert.Equal(t, 1, len(order.Items))
		assert.Equal(t, 1, order.Items[0].Quantity)
		assert.Equal(t, 10.00, order.Total())
	})

	t.Run("get the total price order whith thow itens", func(t *testing.T) {
		order.AddItem(item2)
		assert.Equal(t, 2, len(order.Items))
		assert.Equal(t, 1, order.Items[0].Quantity)
		assert.Equal(t, 3, order.Items[1].Quantity)
		assert.Equal(t, 40.00, order.Total())
	})

	t.Run("get the display total price", func(t *testing.T) {
		assert.Equal(t, "40.00", order.DisplayTotalPrice())
	})
}
