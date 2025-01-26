package cart

import (
	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/fail"
	"github.com/Marlliton/validator/rule"
)

type Item struct {
	ID        id.ID
	CartID    id.ID
	ProductID id.ID
	AddOns    []*AddOn
	Price     int
	Quantity  int
}

func NewItem(cartID, prodID id.ID, price, quantity int) (*Item, []*fail.Error) {
	ci := &Item{
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

func (it *Item) GetTotalPrice() int {
	addOnsTotal := 0
	for _, ad := range it.AddOns {
		addOnsTotal += ad.GetTotalPrice()
	}
	return (it.Price * it.Quantity) + addOnsTotal
}

func (it *Item) IncludeAddOn(addOn *AddOn) {
	for i, ad := range it.AddOns {
		if ad.ID == addOn.ID {
			it.AddOns[i].Quantity = addOn.Quantity
			return
		}
	}

	it.AddOns = append(it.AddOns, addOn)
}

func (it *Item) RemoveAddOn(id id.ID) {
	for i, ad := range it.AddOns {
		if ad.ID == id {
			it.AddOns = append(it.AddOns[:i], it.AddOns[i+1:]...)
			return
		}
	}
}

func (ci *Item) validate() (bool, []*fail.Error) {
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
