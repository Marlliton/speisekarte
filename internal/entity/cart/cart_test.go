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
