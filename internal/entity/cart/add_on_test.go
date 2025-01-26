package cart

import (
	"testing"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestAddOn_New(t *testing.T) {
	orderItemID := id.New()
	name := "Extra Cheese"
	price := 200

	t.Run("should create an add-on", func(t *testing.T) {
		addOn, errs := NewAddOn(orderItemID, name, price, 1)

		assert.NotNil(t, addOn)
		assert.Nil(t, errs)
		assert.Equal(t, orderItemID, addOn.CartItemID)
		assert.Equal(t, name, addOn.Name)
		assert.Equal(t, price, addOn.Price)
		assert.Equal(t, 1, addOn.Quantity)
	})

	t.Run("should create an add-on with minimum price and quantity", func(t *testing.T) {
		addOn, errs := NewAddOn(orderItemID, name, 0, 1)

		assert.NotNil(t, addOn)
		assert.Nil(t, errs)
		assert.Equal(t, orderItemID, addOn.CartItemID)
		assert.Equal(t, name, addOn.Name)
		assert.Equal(t, 0, addOn.Price)
		assert.Equal(t, 1, addOn.Quantity)
	})
}

func TestAddOn_GetTotalPrice(t *testing.T) {
	orderItemID := id.New()
	name := "Extra Cheese"
	price := 200
	qty := 4
	total := 800

	t.Run("should get a total price", func(t *testing.T) {
		addOn, errs := NewAddOn(orderItemID, name, price, qty)

		assert.NotNil(t, addOn)
		assert.Nil(t, errs)
		assert.Equal(t, total, addOn.GetTotalPrice())
	})
}

func TestAddOn_FailToCreate(t *testing.T) {
	orderItemID := id.New()
	t.Run("should fail to create without OrderItemID", func(t *testing.T) {
		addOn, errs := NewAddOn(id.ID{}, "Extra Cheese", 100, 1)

		assert.Nil(t, addOn)
		assert.NotNil(t, errs)
	})

	t.Run("should failt to create without name", func(t *testing.T) {
		addOn, errs := NewAddOn(orderItemID, "", 100, 1)

		assert.Nil(t, addOn)
		assert.NotNil(t, errs)
	})

	t.Run("should failt to create with a short name", func(t *testing.T) {
		addOn, errs := NewAddOn(orderItemID, "Ex", 100, 1)

		assert.Nil(t, addOn)
		assert.NotNil(t, errs)
	})

	t.Run("should failt to create with a long name", func(t *testing.T) {
		addOn, errs := NewAddOn(orderItemID, "This name is way too long for validation", 100, 1)

		assert.Nil(t, addOn)
		assert.NotNil(t, errs)
	})

	t.Run("should failt to create with negative price", func(t *testing.T) {
		addOn, errs := NewAddOn(orderItemID, "Ex", -100, 1)

		assert.Nil(t, addOn)
		assert.NotNil(t, errs)
	})

	t.Run("should failt to create with negative invalid quantity", func(t *testing.T) {
		addOn, errs := NewAddOn(orderItemID, "Ex", 100, 0)

		assert.Nil(t, addOn)
		assert.NotNil(t, errs)
	})
}
