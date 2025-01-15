package order

import (
	"time"

	"github.com/Marlliton/speisekarte/pkg/id"
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
	Products   []OrderProduct
	Status     OrderStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type OrderProduct struct {
	ProductID id.ID
	Quantity  int
}
