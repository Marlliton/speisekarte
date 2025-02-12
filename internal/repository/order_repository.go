package repository

import (
	"context"

	"github.com/Marlliton/speisekarte/internal/entity/order"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/Marlliton/speisekarte/pkg/id"
)

type OrderRepository interface {
	Create(ctx context.Context, order *order.Order) *apperr.AppErr
	FindAllByCustomerID(customerID id.ID) ([]*order.Order, *apperr.AppErr)
}
