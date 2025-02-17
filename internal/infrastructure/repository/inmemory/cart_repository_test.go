package inmemory

import (
	"context"
	"testing"

	"github.com/Marlliton/speisekarte/internal/entity/cart"
	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestCartRepository(t *testing.T) {
	var repo *cartRepository
	ctx := context.Background()

	setup := func() {
		repo = NewInMemoryCartRepository()
	}

	t.Run("should create a new cart", func(t *testing.T) {
		setup()
		cartID := id.New()
		customerID := id.New()
		c := &cart.Cart{ID: cartID, CustomerID: customerID}
		err := repo.Create(ctx, c)
		assert.Nil(t, err)

		found, err := repo.FindById(ctx, cartID)
		assert.Nil(t, err)
		assert.Equal(t, cartID, found.ID)
	})

	t.Run("should return error when cart is not found", func(t *testing.T) {
		setup()
		_, err := repo.FindById(ctx, id.New())
		assert.NotNil(t, err)
	})

	t.Run("should find cart by customer ID", func(t *testing.T) {
		setup()
		cartID := id.New()
		customerID := id.New()
		c := &cart.Cart{ID: cartID, CustomerID: customerID}
		repo.Create(ctx, c)

		found, err := repo.FindByCustomerId(ctx, customerID)
		assert.Nil(t, err)
		assert.Equal(t, customerID, found.CustomerID)
	})

	t.Run("should update a cart", func(t *testing.T) {
		setup()
		cartID := id.New()
		customerID := id.New()
		c := &cart.Cart{ID: cartID, CustomerID: customerID}
		repo.Create(ctx, c)

		updatedCart := &cart.Cart{ID: cartID, CustomerID: customerID, Items: []*cart.Item{}}
		err := repo.Update(ctx, updatedCart)
		assert.Nil(t, err)

		found, err := repo.FindById(ctx, cartID)
		assert.Nil(t, err)
		assert.Equal(t, 0, len(found.Items))
	})

	t.Run("should delete a cart", func(t *testing.T) {
		setup()
		cartID := id.New()
		customerID := id.New()
		c := &cart.Cart{ID: cartID, CustomerID: customerID}
		repo.Create(ctx, c)

		err := repo.Delete(ctx, cartID)
		assert.Nil(t, err)

		_, err = repo.FindById(ctx, cartID)
		assert.NotNil(t, err)
	})

	t.Run("should add an item to the cart", func(t *testing.T) {
		setup()
		cartID := id.New()
		customerID := id.New()
		c := &cart.Cart{ID: cartID, CustomerID: customerID}
		repo.Create(ctx, c)

		item, errs := cart.NewItem(cartID, id.New(), 1000, 1)
		assert.Empty(t, errs)
		err := repo.AddItem(ctx, cartID, item)
		assert.Nil(t, err)

		found, err := repo.FindById(ctx, cartID)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(found.Items))
	})

	t.Run("should update an item in the cart", func(t *testing.T) {
		setup()
		cartID := id.New()
		customerID := id.New()
		c := &cart.Cart{ID: cartID, CustomerID: customerID}
		repo.Create(ctx, c)

		item, errs := cart.NewItem(cartID, id.New(), 1000, 1)
		assert.Empty(t, errs)
		repo.AddItem(ctx, cartID, item)

		t.Logf("Cart: %+v\n", c.Items)

		updatedItem := &cart.Item{
			CartID:    c.ID,
			ID:        item.ID,
			ProductID: item.ProductID,
			Price:     1200,
			Quantity:  1,
			AddOns:    item.AddOns,
		}
		assert.Empty(t, errs)
		err := repo.UpdateItem(ctx, cartID, updatedItem)
		assert.Nil(t, err)

		found, err := repo.FindById(ctx, cartID)
		assert.Nil(t, err)
		assert.Equal(t, 1200, found.Items[0].Price)
	})

	t.Run("should remove an item from the cart", func(t *testing.T) {
		setup()
		cartID := id.New()
		customerID := id.New()
		c := &cart.Cart{ID: cartID, CustomerID: customerID}
		repo.Create(ctx, c)

		item, errs := cart.NewItem(cartID, id.New(), 1000, 1)
		assert.Empty(t, errs)
		repo.AddItem(ctx, cartID, item)

		err := repo.RemoveItem(ctx, cartID, item.ID)
		assert.Nil(t, err)

		found, err := repo.FindById(ctx, cartID)
		assert.Nil(t, err)
		assert.Equal(t, 0, len(found.Items))
	})

	t.Run("should add an add-on to an item in the cart", func(t *testing.T) {
		setup()
		cartID := id.New()
		customerID := id.New()
		c := &cart.Cart{ID: cartID, CustomerID: customerID}
		repo.Create(ctx, c)

		item, errs := cart.NewItem(cartID, id.New(), 1000, 1)
		assert.Empty(t, errs)
		repo.AddItem(ctx, cartID, item)

		addOn, errs := cart.NewAddOn(item.ID, "Extra Cheese", 200, 1)
		assert.Empty(t, errs)
		err := repo.AddAddOnToItem(ctx, cartID, item.ID, addOn)
		assert.Nil(t, err)

		found, err := repo.FindById(ctx, cartID)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(found.Items[0].AddOns))
	})

	t.Run("should update an add-on in an item in the cart", func(t *testing.T) {
		setup()
		cartID := id.New()
		customerID := id.New()
		c := &cart.Cart{ID: cartID, CustomerID: customerID}
		repo.Create(ctx, c)

		item, errs := cart.NewItem(cartID, id.New(), 1000, 1)
		assert.Empty(t, errs)
		repo.AddItem(ctx, cartID, item)

		addOn, errs := cart.NewAddOn(item.ID, "Extra Cheese", 200, 1)
		assert.Empty(t, errs)
		repo.AddAddOnToItem(ctx, cartID, item.ID, addOn)

		updatedAddOn := &cart.AddOn{
			ID:       addOn.ID,
			ItemID:   addOn.ItemID,
			Name:     "Extra Cheese",
			Price:    250, // Novo preço
			Quantity: 1,
		}

		assert.Empty(t, errs)
		err := repo.UpdateAddOnInItem(ctx, cartID, item.ID, updatedAddOn)
		assert.Nil(t, err)

		found, err := repo.FindById(ctx, cartID)
		assert.Nil(t, err)
		assert.Equal(t, 250, found.Items[0].AddOns[0].Price)
	})

	t.Run("should remove an add-on from an item in the cart", func(t *testing.T) {
		setup()
		cartID := id.New()
		customerID := id.New()
		c := &cart.Cart{ID: cartID, CustomerID: customerID}
		repo.Create(ctx, c)

		item, errs := cart.NewItem(cartID, id.New(), 1000, 1)
		assert.Empty(t, errs)
		repo.AddItem(ctx, cartID, item)

		addOn, errs := cart.NewAddOn(item.ID, "Extra Cheese", 200, 1)
		assert.Empty(t, errs)
		repo.AddAddOnToItem(ctx, cartID, item.ID, addOn)

		err := repo.RemoveAddOnFromItem(ctx, cartID, item.ID, addOn.ID)
		assert.Nil(t, err)

		found, err := repo.FindById(ctx, cartID)
		assert.Nil(t, err)
		assert.Equal(t, 0, len(found.Items[0].AddOns))
	})
}
