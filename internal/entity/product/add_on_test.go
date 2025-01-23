package product

import (
	"testing"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestAddOn_New(t *testing.T) {
	productID := id.New()
	name := "Extra Cheese"
	price := 100

	t.Run("should create an add-on", func(t *testing.T) {
		addOn, errs := NewAddOn(productID, name, price)

		assert.Nil(t, errs)
		assert.NotNil(t, addOn)
		assert.Equal(t, productID, addOn.ProductID)
		assert.Equal(t, name, addOn.Name)
		assert.Equal(t, price, addOn.Price)
	})

	t.Run("should create an add-on with minimum price", func(t *testing.T) {
		addOn, errs := NewAddOn(productID, name, 0)

		assert.Nil(t, errs)
		assert.NotNil(t, addOn)
		assert.Equal(t, productID, addOn.ProductID)
		assert.Equal(t, name, addOn.Name)
		assert.Equal(t, 0, addOn.Price)
	})
}

func TestAddOn_FailToCreate(t *testing.T) {
	productID := id.New()

	t.Run("should fail to create without ProductID", func(t *testing.T) {
		addOn, errs := NewAddOn(id.ID{}, "Extra Cheese", 100)

		assert.NotNil(t, errs)
		assert.Nil(t, addOn)
	})

	t.Run("should fail to create without Name", func(t *testing.T) {
		addOn, errs := NewAddOn(productID, "", 100)

		assert.NotNil(t, errs)
		assert.Nil(t, addOn)
	})

	t.Run("should fail to create with a short Name", func(t *testing.T) {
		addOn, errs := NewAddOn(productID, "Ex", 100)

		assert.NotNil(t, errs)
		assert.Nil(t, addOn)
	})

	t.Run("should fail to create with a long Name", func(t *testing.T) {
		addOn, errs := NewAddOn(productID, "ThisNameIsWayTooLongForValidation", 100)

		assert.NotNil(t, errs)
		assert.Nil(t, addOn)
	})

	t.Run("should fail to create with negative Price", func(t *testing.T) {
		addOn, errs := NewAddOn(productID, "Extra Cheese", -10)

		assert.NotNil(t, errs)
		assert.Nil(t, addOn)
	})
}
