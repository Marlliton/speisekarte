package order

import (
	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/fail"
	"github.com/Marlliton/validator/rule"
)

type OrderItem struct {
	ID        id.ID
	OrderID   id.ID
	ProductID id.ID
	AddOns    []*OrderItemAddOn
	Quantity  int
	Price     int
}

func OrderNewItem(
	orderID, productID id.ID, quantity, price int, addOns ...*OrderItemAddOn,
) (*OrderItem, []*fail.Error) {
	if len(addOns) == 0 {
		addOns = make([]*OrderItemAddOn, 0)
	}
	o := &OrderItem{ID: id.New(),
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

func (o *OrderItem) validate() (bool, []*fail.Error) {
	v := validator.New()
	v.Add("OrderID", rule.Rules{rule.Required()})
	v.Add("ProductID", rule.Rules{rule.Required()})
	v.Add("Quantity", rule.Rules{rule.Required(), rule.MinValue(1)})
	v.Add("Price", rule.Rules{rule.Required(), rule.MinValue(0)})

	if errs := v.Validate(*o); len(errs) > 0 {
		return false, errs
	}

	return true, nil
}
