package cart

import (
	"fmt"
	"time"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/fail"
	"github.com/Marlliton/validator/rule"
)

type priceDisplay struct {
	Total    string
	SubTotal string
}

type Cart struct {
	ID         id.ID
	CustomerID id.ID
	Items      []*Item
	Rate       int
	Discount   int
	Total      int
	SubTotal   int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func New(customerID id.ID, rate, discount int, items ...*Item) (*Cart, []*fail.Error) {
	if len(items) == 0 {
		items = make([]*Item, 0)
	}
	c := &Cart{
		ID:         id.New(),
		CustomerID: customerID,
		Rate:       rate,
		Discount:   discount,
		Items:      items,
		CreatedAt:  time.Now(),
	}
	total := c.calculateTotal()
	subTotal := c.calculateSubTotal()
	c.Total = total
	c.SubTotal = subTotal

	if ok, errs := c.validate(); !ok {
		return nil, errs
	}

	return c, nil
}

// NOTE: Valor total da compra
func (c *Cart) calculateSubTotal() int {
	subTotal := 0
	for _, item := range c.Items {
		subTotal += item.GetTotalPrice()
	}
	return subTotal
}

// NOTE: Valor final que o cliente tem que pagar
func (c *Cart) calculateTotal() int {
	subTotal := c.calculateSubTotal()

	total := subTotal + c.Rate - c.Discount

	return total
}

func (c *Cart) DisplayTotalPrice() priceDisplay {
	t := c.calculateTotal()
	subT := c.calculateSubTotal()

	total := fmt.Sprintf("%.2f", float64(t)/100)
	subTotal := fmt.Sprintf("%.2f", float64(subT)/100)

	return priceDisplay{total, subTotal}
}

func (c *Cart) AddItem(item *Item) {
	for i, it := range c.Items {
		if it.ProductID == item.ProductID {
			c.Items[i].Quantity = item.Quantity
			c.UpdatedAt = time.Now()
			return
		}
	}

	c.Items = append(c.Items, item) // NOTE: Se o item não existir adiciona no cart
	c.UpdatedAt = time.Now()
}

func (c *Cart) RemoveItem(id id.ID) {
	for i, it := range c.Items {
		if it.ID == id {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			c.UpdatedAt = time.Now()
		}
	}
}

func (c *Cart) validate() (bool, []*fail.Error) {
	v := validator.New()
	v.Add("ID", rule.Rules{rule.Required()})
	v.Add("CustomerID", rule.Rules{rule.Required()})
	v.Add("Items", rule.Rules{rule.MinLength(0)})
	v.Add("Rate", rule.Rules{rule.Required(), rule.MinValue(0)})
	v.Add("Discount", rule.Rules{rule.Required(), rule.MinValue(0)})
	v.Add("Total", rule.Rules{rule.Required(), rule.MinValue(0)})
	v.Add("SubTotal", rule.Rules{rule.Required(), rule.MinValue(0)})
	v.Add("CreatedAt", rule.Rules{rule.Required()})

	if errs := v.Validate(*c); len(errs) > 0 {
		return false, errs
	}

	return true, nil
}
