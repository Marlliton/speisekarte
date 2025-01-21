package product

import (
	"time"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/fail"
	"github.com/Marlliton/validator/rule"
)

type AddOn struct {
	ID        id.ID
	ProductID id.ID
	Name      string
	Price     int
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAddOn(productID id.ID, name string, price, quantity int) (*AddOn, []*fail.Error) {
	addOn := &AddOn{
		ID:        id.New(),
		ProductID: productID,
		Name:      name,
		Price:     price,
		Quantity:  quantity,
		CreatedAt: time.Now(),
	}

	if ok, errs := addOn.validate(); !ok {
		return nil, errs
	}

	return addOn, nil
}

func (a *AddOn) validate() (bool, []*fail.Error) {
	v := validator.New()
	v.Add("ProductID", rule.Rules{rule.Required()})
	v.Add("Name", rule.Rules{rule.Required(),
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
