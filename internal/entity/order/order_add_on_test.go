package order

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
		addOn, errs := OrderItemNewAddOn(orderItemID, name, price, 1)

		assert.NotNil(t, addOn)
		assert.Nil(t, errs)
		assert.Equal(t, orderItemID, addOn.OrderItemID)
		assert.Equal(t, name, addOn.Name)
		assert.Equal(t, price, addOn.Price)
		assert.Equal(t, 1, addOn.Quantity)
	})

	t.Run("should create an add-on with minimum price and quantity", func(t *testing.T) {
		addOn, errs := OrderItemNewAddOn(orderItemID, name, 0, 1)

		assert.NotNil(t, addOn)
		assert.Nil(t, errs)
		assert.Equal(t, orderItemID, addOn.OrderItemID)
		assert.Equal(t, name, addOn.Name)
		assert.Equal(t, 0, addOn.Price)
		assert.Equal(t, 1, addOn.Quantity)
	})
}

func TestAddOn_FailToCreate(t *testing.T) {
	orderItemID := id.New()
	t.Run("should fail to create without OrderItemID", func(t *testing.T) {
		addOn, errs := OrderItemNewAddOn(id.ID{}, "Extra Cheese", 100, 1)

		assert.Nil(t, addOn)
		assert.NotNil(t, errs)
	})

	t.Run("should failt to create without name", func(t *testing.T) {
		addOn, errs := OrderItemNewAddOn(orderItemID, "", 100, 1)

		assert.Nil(t, addOn)
		assert.NotNil(t, errs)
	})

	t.Run("should failt to create with a short name", func(t *testing.T) {
		addOn, errs := OrderItemNewAddOn(orderItemID, "Ex", 100, 1)

		assert.Nil(t, addOn)
		assert.NotNil(t, errs)
	})

	t.Run("should failt to create with a long name", func(t *testing.T) {
		addOn, errs := OrderItemNewAddOn(orderItemID, "This name is way too long for validation", 100, 1)

		assert.Nil(t, addOn)
		assert.NotNil(t, errs)
	})

	t.Run("should failt to create with negative price", func(t *testing.T) {
		addOn, errs := OrderItemNewAddOn(orderItemID, "Ex", -100, 1)

		assert.Nil(t, addOn)
		assert.NotNil(t, errs)
	})

	t.Run("should failt to create with negative invalid quantity", func(t *testing.T) {
		addOn, errs := OrderItemNewAddOn(orderItemID, "Ex", 100, 0)

		assert.Nil(t, addOn)
		assert.NotNil(t, errs)
	})
}
