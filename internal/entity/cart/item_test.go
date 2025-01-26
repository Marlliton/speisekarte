package cart

import (
	"testing"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestNewCartItem_Creation(t *testing.T) {
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

func TestCartItem_IncludeAddOn(t *testing.T) {
	cartID := id.New()
	prodID := id.New()
	price := 900
	quantity := 2

	t.Run("should include an add-on to the item", func(t *testing.T) {
		item, errs := NewItem(cartID, prodID, price, quantity)
		assert.Nil(t, errs)
		assert.NotNil(t, item)

		addOn1, errs := NewAddOn(cartID, "Extra Cheese", 500, 2)
		assert.Nil(t, errs)
		assert.NotNil(t, addOn1)

		item.IncludeAddOn(addOn1)
		assert.Len(t, item.AddOns, 1)
		assert.Equal(t, addOn1.ID, item.AddOns[0].ID)
		assert.Equal(t, addOn1.Quantity, item.AddOns[0].Quantity)

		addOn2, errs := NewAddOn(cartID, "French fries", 1500, 4)
		assert.Nil(t, errs)
		assert.NotNil(t, addOn1)

		item.IncludeAddOn(addOn2)
		assert.Len(t, item.AddOns, 2)
		assert.Equal(t, addOn2.ID, item.AddOns[1].ID)
		assert.Equal(t, addOn2.Quantity, item.AddOns[1].Quantity)
	})

	t.Run("should remove an add-on from the item", func(t *testing.T) {
		item, errs := NewItem(cartID, prodID, price, quantity)
		assert.Nil(t, errs)
		assert.NotNil(t, item)

		addOn1, errs := NewAddOn(cartID, "Extra Cheese", 500, 2)
		addOn2, errs := NewAddOn(cartID, "French fries", 1500, 4)
		assert.Nil(t, errs)
		assert.NotNil(t, addOn1)

		item.IncludeAddOn(addOn1)
		item.IncludeAddOn(addOn2)
		assert.Len(t, item.AddOns, 2)
		assert.Equal(t, addOn2.ID, item.AddOns[1].ID)
		assert.Equal(t, addOn2.Quantity, item.AddOns[1].Quantity)

		// remove addon
		item.RemoveAddOn(addOn2.ID)
		assert.Len(t, item.AddOns, 1)
		assert.NotNil(t, item.AddOns[0])
		assert.Len(t, item.AddOns, 1)
		assert.Equal(t, addOn1.ID, item.AddOns[0].ID)
	})
}

func TestItem_GetTotalPrice(t *testing.T) {
	orderItemID := id.New()
	name := "Extra Cheese"
	price := 200
	qty := 4

	cartID := id.New()
	prodID := id.New()
	itemPrice := 900
	quantity := 2
	itemTotal := 2600

	t.Run("should get a total price", func(t *testing.T) {
		addOn, errs := NewAddOn(orderItemID, name, price, qty)
		assert.Nil(t, errs)
		assert.NotNil(t, addOn)

		item, errs := NewItem(cartID, prodID, itemPrice, quantity)
		assert.Nil(t, errs)
		assert.NotNil(t, item)
		item.IncludeAddOn(addOn)

		assert.Equal(t, itemTotal, item.GetTotalPrice())
	})
}

func TestCartItem_Validation(t *testing.T) {
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
