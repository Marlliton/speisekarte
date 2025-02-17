package repository

import (
	"context"

	"github.com/Marlliton/speisekarte/internal/entity/customer"
	"github.com/Marlliton/speisekarte/pkg/apperr"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer *customer.Customer) *apperr.AppErr
	FindByPhone(ctx context.Context, phone string) (*customer.Customer, *apperr.AppErr)
}
