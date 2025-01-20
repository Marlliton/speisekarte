package cart

import (
	"testing"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func createCart(t *testing.T, customerID id.ID, fee, discount int, items ...*CartItem) *Cart {
	t.Helper()
	c, errs := New(customerID, fee, discount, items...)
	assert.Nil(t, errs)

	return c
}

func createCartItem(price, quantity int) *CartItem {
	return &CartItem{
		ID:        id.New(),
		ProductID: id.New(),
		Price:     price,
		Quantity:  quantity,
	}

}

func TestNewCart_Success(t *testing.T) {
	customerID := id.New()
	cartItem := []*CartItem{
		{
			ID:        id.New(),
			ProductID: id.New(),
			Price:     100,
			Quantity:  3,
		},
	}

	c, errs := New(customerID, 2, 1, cartItem...)
	assert.Nil(t, errs)
	assert.NotNil(t, c)
	assert.Len(t, c.Items, 1)
}

func TestCart_Totals(t *testing.T) {
	priceInCents := 2000 // NOTE: 20 Reais
	qty := 2
	fee := 100      // NOTE: Delivery fee in cents 1 real
	discount := 400 // NOTE: Discount in cents 4 reais
	t.Run("get the total cart", func(t *testing.T) {
		it := createCartItem(priceInCents, qty)
		c := createCart(t, id.New(), fee, discount, it)

		expectedTotalInCents := 3700
		assert.Equal(t, expectedTotalInCents, c.Total)
	})

	t.Run("get the sub total cart", func(t *testing.T) {
		it := createCartItem(priceInCents, qty)
		c := createCart(t, id.New(), fee, discount, it)

		expectedTotalInCents := 4000
		assert.Equal(t, expectedTotalInCents, c.SubTotal)
	})

	t.Run("get the display total and sub total", func(t *testing.T) {
		it := createCartItem(priceInCents, qty)
		c := createCart(t, id.New(), fee, discount, it)

		expectedTotalInCents := "37.00"
		expectedSubTotalInCents := "40.00"

		pDisplay := c.DisplayTotalPrice()
		assert.Equal(t, expectedTotalInCents, pDisplay.Total)
		assert.Equal(t, expectedSubTotalInCents, pDisplay.SubTotal)
	})
}

// TODO: teste dos itens
func TestCart_AddAndRemoveItems(t *testing.T) {
	t.Run("Add a new item", func(t *testing.T) {
		it := createCartItem(800, 1)
		c := createCart(t, id.New(), 0, 0)
		assert.NotNil(t, c)
		assert.Len(t, c.Items, 0)

		c.AddItem(it)
		assert.Len(t, c.Items, 1)
		assert.Equal(t, 1, c.Items[0].Quantity)
	})

	t.Run("Add and modify quantity of item", func(t *testing.T) {
		it := createCartItem(800, 1)
		c := createCart(t, id.New(), 0, 0)
		c.AddItem(it)
		assert.Len(t, c.Items, 1)
		assert.Equal(t, 1, c.Items[0].Quantity)

		it.Quantity = 4
		c.AddItem(it)
		assert.Len(t, c.Items, 1)
		assert.Equal(t, 4, c.Items[0].Quantity)

		it2 := createCartItem(200, 2)
		c.AddItem(it2)
		assert.Len(t, c.Items, 2)
		assert.Equal(t, 4, c.Items[0].Quantity)
		assert.Equal(t, 2, c.Items[1].Quantity)
	})

	t.Run("Remove item from cart", func(t *testing.T) {
		it := createCartItem(800, 1)
		it2 := createCartItem(800, 1)
		c := createCart(t, id.New(), 0, 0, it, it2)

		assert.Len(t, c.Items, 2)

		c.RemoveItem(it2)
		assert.Len(t, c.Items, 1)
	})
}
