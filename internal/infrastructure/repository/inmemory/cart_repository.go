package inmemory

import (
	"context"
	"sync"

	"github.com/Marlliton/speisekarte/internal/entity/cart"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/Marlliton/speisekarte/pkg/id"
)

type cartRepository struct {
	sync.RWMutex
	carts map[id.ID]*cart.Cart
}

func NewInMemoryCartRepository() *cartRepository {
	return &cartRepository{
		carts: make(map[id.ID]*cart.Cart, 0),
	}
}

func (r *cartRepository) Create(ctx context.Context, cart *cart.Cart) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.carts[cart.ID]; exists {
		return apperr.New("already exists")
	}
	r.carts[cart.ID] = cart

	return nil
}

func (r *cartRepository) FindById(ctx context.Context, cartID id.ID) (*cart.Cart, *apperr.AppErr) {
	r.RLock()
	defer r.RUnlock()

	if _, exists := r.carts[cartID]; !exists {
		return nil, apperr.New("not found")
	}

	return r.carts[cartID], nil
}

func (r *cartRepository) FindByCustomerId(ctx context.Context, customerID id.ID) (*cart.Cart, *apperr.AppErr) {
	r.RLock()
	defer r.RUnlock()

	for _, cart := range r.carts {
		if cart.CustomerID == customerID {
			return cart, nil
		}

	}

	return nil, apperr.New("cart not found for customer")
}

func (r *cartRepository) Update(ctx context.Context, cart *cart.Cart) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.carts[cart.ID]; !exists {
		return apperr.New("car not found")
	}

	r.carts[cart.ID] = cart
	return nil
}

func (r *cartRepository) Delete(ctx context.Context, cartID id.ID) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.carts[cartID]; !exists {
		return apperr.New("car not found")
	}

	delete(r.carts, cartID)
	return nil
}

func (r *cartRepository) AddItem(ctx context.Context, cartID id.ID, item *cart.Item) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()
	cart, exists := r.carts[cartID]
	if !exists {
		return apperr.New("car not found")
	}

	cart.Items = append(cart.Items, item)
	return nil
}

func (r *cartRepository) UpdateItem(ctx context.Context, cartID id.ID, item *cart.Item) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	cart, exists := r.carts[cartID]
	if !exists {
		return apperr.New("car not found")
	}

	for i, it := range cart.Items {
		if it.ID == item.ID {
			cart.Items[i] = item
			return nil
		}
	}

	return apperr.New("item not found in cart")
}
func (r *cartRepository) RemoveItem(ctx context.Context, cartID id.ID, itemID id.ID) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	cart, exists := r.carts[cartID]
	if !exists {
		return apperr.New("car not found")
	}

	for i, it := range cart.Items {
		if it.ID == itemID {
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			return nil
		}
	}

	return nil
}

func (r *cartRepository) AddAddOnToItem(ctx context.Context, cartID id.ID, itemID id.ID, addOn *cart.AddOn) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	cart, exists := r.carts[cartID]
	if !exists {
		return apperr.New("car not found")
	}

	for i, it := range cart.Items {
		if it.ID == itemID {
			cart.Items[i].AddOns = append(cart.Items[i].AddOns, addOn)
			return nil
		}
	}

	return apperr.New("item not found in cart")
}
func (r *cartRepository) UpdateAddOnInItem(ctx context.Context, cartID id.ID, itemID id.ID, addOn *cart.AddOn) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	cart, exists := r.carts[cartID]
	if !exists {
		return apperr.New("car not found")
	}

	for _, it := range cart.Items {
		if it.ID == itemID {
			for i, ad := range it.AddOns {
				if ad.ID == addOn.ID {
					it.AddOns[i] = addOn
					return nil
				}
			}
		}
	}

	return nil
}
func (r *cartRepository) RemoveAddOnFromItem(ctx context.Context, cartID id.ID, itemID id.ID, addOnID id.ID) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	cart, exists := r.carts[cartID]
	if !exists {
		return apperr.New("car not found")
	}

	for _, it := range cart.Items {
		if it.ID == itemID {
			for i, ad := range it.AddOns {
				if ad.ID == addOnID {
					it.AddOns = append(it.AddOns[:i], it.AddOns[i+1:]...)
					return nil
				}
			}
		}
	}

	return nil
}
