package repository

import (
	"context"

	"github.com/Marlliton/speisekarte/internal/entity/cart"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/Marlliton/speisekarte/pkg/id"
)

type CartRepository interface {
	Create(ctx context.Context, cart *cart.Cart) *apperr.AppErr
	FindById(ctx context.Context, cartID id.ID) (*cart.Cart, *apperr.AppErr)
	FindByCustomerId(ctx context.Context, customerID id.ID) (*cart.Cart, *apperr.AppErr)
	Update(ctx context.Context, cart *cart.Cart) *apperr.AppErr
	Delete(ctx context.Context, cartID id.ID) *apperr.AppErr

	AddItem(ctx context.Context, cartID id.ID, item *cart.Item) *apperr.AppErr
	UpdateItem(ctx context.Context, cartID id.ID, item *cart.Item) *apperr.AppErr
	RemoveItem(ctx context.Context, cartID id.ID, itemID id.ID) *apperr.AppErr

	AddAddOnToItem(ctx context.Context, cartID id.ID, itemID id.ID, addOn *cart.AddOn) *apperr.AppErr
	UpdateAddOnInItem(ctx context.Context, cartID id.ID, itemID id.ID, addOn *cart.AddOn) *apperr.AppErr
	RemoveAddOnFromItem(ctx context.Context, cartID id.ID, itemID id.ID, addOnID id.ID) *apperr.AppErr
}
