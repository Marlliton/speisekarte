package cart

import (
	"fmt"
	"time"

	"github.com/Marlliton/speisekarte/pkg/id"
)

type total string
type subTotal string

type Cart struct {
	ID          id.ID
	CustomerID  id.ID
	Items       []*CartItem
	DeliveryFee int
	Discount    int
	Total       int
	SubTotal    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CartItem struct {
	ProductID id.ID
	Price     int
	Quantity  int
}

func (c *Cart) calculateSubTotal() int {
	subTotal := 0
	for _, item := range c.Items {
		subTotal += item.Price * item.Quantity
	}
	return subTotal
}

func (c *Cart) calculateTotal() int {
	subTotal := c.calculateTotal()

	total := subTotal + c.DeliveryFee - c.Discount

	return total
}

func (c *Cart) DisplayTotalPrice() (total, subTotal) {
	t := c.calculateTotal()
	subT := c.calculateSubTotal()

	total := total(fmt.Sprintf("%.2f", float64(t)/100))
	subTotal := subTotal(fmt.Sprintf("%.2f", float64(subT)/100))

	return total, subTotal
}

func (c *Cart) AddItem(item *CartItem) {
	for i, it := range c.Items {
		if it.ProductID == item.ProductID {
			c.Items[i].Quantity = item.Quantity
			return
		}
	}

	c.Items = append(c.Items, item)
	c.UpdatedAt = time.Now()
}
