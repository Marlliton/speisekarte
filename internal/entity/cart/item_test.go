package cart

import (
	"testing"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestNewItem_Creation(t *testing.T) {
	t.Run("should create a new item successfully", func(t *testing.T) {
		cartID := id.New()
		prodID := id.New()
		price := 900
		quantity := 2

		item, errs := NewItem(cartID, prodID, price, quantity)

		assert.Nil(t, errs)
		assert.NotNil(t, item)
		assert.Equal(t, cartID, item.CartID)
		assert.Equal(t, prodID, item.ProductID)
		assert.Equal(t, price, item.Price)
		assert.Equal(t, quantity, item.Quantity)
	})
}

func TestNewItem_Validation(t *testing.T) {
	t.Run("should fail when CartID is missing", func(t *testing.T) {
		prodID := id.New()
		price := 900
		quantity := 1

		item, errs := NewItem(id.ID{}, prodID, price, quantity)

		assert.Nil(t, item)
		assert.NotNil(t, errs)
		assert.Len(t, errs, 1)
	})

	t.Run("should fail when ProductID is missing", func(t *testing.T) {
		cartID := id.New()
		price := 900
		quantity := 1

		item, errs := NewItem(cartID, id.ID{}, price, quantity)

		assert.Nil(t, item)
		assert.NotNil(t, errs)
		assert.Len(t, errs, 1)
	})

	t.Run("should fail when Price is negative", func(t *testing.T) {
		cartID := id.New()
		prodID := id.New()
		price := -100
		quantity := 1

		item, errs := NewItem(cartID, prodID, price, quantity)

		assert.Nil(t, item)
		assert.NotNil(t, errs)
		assert.Len(t, errs, 1)
	})

	t.Run("should fail when Quantity is zero", func(t *testing.T) {
		cartID := id.New()
		prodID := id.New()
		price := 900
		quantity := 0

		item, errs := NewItem(cartID, prodID, price, quantity)

		assert.Nil(t, item)
		assert.NotNil(t, errs)
		assert.Len(t, errs, 1)
	})
}
