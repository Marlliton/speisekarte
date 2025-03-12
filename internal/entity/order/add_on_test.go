package order

import (
	"testing"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestNewAddOn_Success(t *testing.T) {
	orderItemID := id.New()

	t.Run("should create an add-on", func(t *testing.T) {
		addOn, errs := NewAddOn(orderItemID, "Extra Cheese", 200, 1)

		assert.Nil(t, errs)
		assert.NotNil(t, addOn)
		assert.Equal(t, orderItemID, addOn.OrderItemID)
		assert.Equal(t, "Extra Cheese", addOn.Name)
		assert.Equal(t, 200, addOn.Price)
		assert.Equal(t, 1, addOn.Quantity)
	})
}

func TestNewAddOn_FailToCreate(t *testing.T) {
	t.Run("should fail to create without order item id", func(t *testing.T) {
		addOn, errs := NewAddOn(id.ID{}, "Extra Cheese", 200, 1)

		assert.Nil(t, addOn)
		assert.NotNil(t, errs)
	})

	t.Run("should fail to create with invalid name", func(t *testing.T) {
		addOn, errs := NewAddOn(id.New(), "Ex", 200, 1)

		assert.Nil(t, addOn)
		assert.NotNil(t, errs)
	})

	t.Run("should fail to create with invalid price", func(t *testing.T) {
		addOn, errs := NewAddOn(id.New(), "Extra Cheese", -200, 1)

		assert.Nil(t, addOn)
		assert.NotNil(t, errs)
	})

	t.Run("should fail to create with invalid quantity", func(t *testing.T) {
		addOn, errs := NewAddOn(id.New(), "Extra Cheese", 200, 0)

		assert.Nil(t, addOn)
		assert.NotNil(t, errs)
	})
}
