package order

import (
	"time"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/fail"
	"github.com/Marlliton/validator/rule"
)

type OrderStatus string

const (
	Pending   OrderStatus = "Pending"
	Confirmed OrderStatus = "Confirmed"
	Delivered OrderStatus = "Delivered"
	Canceled  OrderStatus = "Canceled"
)

type Order struct {
	ID         id.ID
	CustomerID id.ID
	Items      []*Item
	Status     OrderStatus
	Total      int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func New(customerID id.ID, items ...*Item) (*Order, []*fail.Error) {
	o := &Order{
		ID:         id.New(),
		CustomerID: customerID,
		Items:      items,
		Status:     Pending,
		CreatedAt:  time.Now(),
	}

	if ok, errs := o.validate(); !ok {
		return nil, errs
	}

	return o, nil
}

func (o *Order) validate() (bool, []*fail.Error) {
	v := validator.New()
	v.Add("CustomerID", rule.Rules{rule.Required()})
	v.Add("Items", rule.Rules{rule.MinLength(1)})
	v.Add("Status", rule.Rules{rule.Required()})

	if errs := v.Validate(*o); len(errs) > 0 {
		return false, errs
	}

	return true, nil
}
