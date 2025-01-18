package order

import (
	"fmt"
	"log"
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
	Items      []*OrderItem
	Status     OrderStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func New(customerID id.ID, items ...*OrderItem) (*Order, []*fail.Error) {
	if items == nil {
		items = []*OrderItem{}
	}
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

func (o *Order) Total() float64 {
	total := 0.0
	for _, item := range o.Items {
		log.Println("aaaaaaaaaaaaaa", item.Quantity)
		total += float64(item.Price) * float64(item.Quantity)
	}
	return total / 100
}

func (o *Order) DisplayTotalPrice() string {
	return fmt.Sprintf("%.2f", o.Total())
}

func (o *Order) validate() (bool, []*fail.Error) {
	v := validator.New()
	v.Add("CustomerID", rule.Rules{rule.Required()})
	v.Add("Status", rule.Rules{rule.Required()})

	if errs := v.Validate(*o); len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

func (o *Order) AddItem(item *OrderItem) {
	for i, it := range o.Items {
		if it.ProductID == item.ProductID {
			o.Items[i].Quantity = item.Quantity
			return
		}
	}

	o.Items = append(o.Items, item)
	o.UpdatedAt = time.Now()
}
