package cart

import (
	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/fail"
	"github.com/Marlliton/validator/rule"
)

type CartItem struct {
	ID        id.ID
	CartID    id.ID
	ProductID id.ID
	Price     int
	Quantity  int
}

func NewCartItem(cartID, prodID id.ID, price, quantity int) (*CartItem, []*fail.Error) {
	ci := &CartItem{
		ID:        id.New(),
		CartID:    cartID,
		ProductID: prodID,
		Price:     price,
		Quantity:  quantity,
	}

	if ok, errs := ci.validate(); !ok {
		return nil, errs
	}

	return ci, nil
}

func (ci *CartItem) validate() (bool, []*fail.Error) {
	v := validator.New()
	v.Add("CartID", rule.Rules{rule.Required()})
	v.Add("ProductID", rule.Rules{rule.Required()})
	v.Add("Price", rule.Rules{rule.Required(), rule.MinValue(0)})
	v.Add("Quantity", rule.Rules{rule.Required(), rule.MinValue(1)})

	if errs := v.Validate(*ci); len(errs) > 0 {
		return false, errs
	}

	return true, nil
}
