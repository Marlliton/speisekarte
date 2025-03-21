package order

import (
	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/fail"
	"github.com/Marlliton/validator/rule"
)

type Item struct {
	ID        id.ID
	OrderID   id.ID
	ProductID id.ID
	AddOns    []*AddOn
	Quantity  int
	Price     int
}

func NewItem(
	id, orderID, productID id.ID, quantity, price int, addOns ...*AddOn,
) (*Item, []*fail.Error) {
	if len(addOns) == 0 {
		addOns = make([]*AddOn, 0)
	}
	o := &Item{
		ID:        id,
		OrderID:   orderID,
		ProductID: productID,
		AddOns:    addOns,
		Quantity:  quantity,
		Price:     price,
	}

	if ok, errs := o.validate(); !ok {
		return nil, errs
	}

	return o, nil
}

func (o *Item) validate() (bool, []*fail.Error) {
	v := validator.New()
	v.Add("ID", rule.Rules{rule.Required()})
	v.Add("OrderID", rule.Rules{rule.Required()})
	v.Add("ProductID", rule.Rules{rule.Required()})
	v.Add("Quantity", rule.Rules{rule.Required(), rule.MinValue(1)})
	v.Add("Price", rule.Rules{rule.Required(), rule.MinValue(0)})

	if errs := v.Validate(*o); len(errs) > 0 {
		return false, errs
	}

	return true, nil
}
