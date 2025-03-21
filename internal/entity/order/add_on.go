package order

import (
	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/fail"
	"github.com/Marlliton/validator/rule"
)

type AddOn struct {
	ID          id.ID
	OrderItemID id.ID
	Name        string
	Price       int
	Quantity    int
}

func NewAddOn(itemID id.ID, name string, price, quantity int) (*AddOn, []*fail.Error) {
	a := &AddOn{
		ID:          id.New(),
		OrderItemID: itemID,
		Name:        name,
		Price:       price,
		Quantity:    quantity,
	}

	if ok, errs := a.validate(); !ok {
		return nil, errs
	}

	return a, nil
}

func (a *AddOn) validate() (bool, []*fail.Error) {
	v := validator.New()
	v.Add("OrderItemID", rule.Rules{rule.Required()})
	v.Add("OrderItemID", rule.Rules{rule.Required()})
	v.Add("Name", rule.Rules{
		rule.Required(),
		rule.MinLength(3),
		rule.MaxLength(20),
	})
	v.Add("Price", rule.Rules{rule.Required(), rule.MinValue(0)})
	v.Add("Quantity", rule.Rules{rule.Required(), rule.MinValue(1)})

	if errs := v.Validate(*a); len(errs) > 0 {
		return false, errs
	}

	return true, nil
}
