package repository

import (
	"context"

	"github.com/Marlliton/speisekarte/internal/entity/customer"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/Marlliton/speisekarte/pkg/id"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer *customer.Customer) *apperr.AppErr
	FindByID(ctx context.Context, id id.ID) (*customer.Customer, *apperr.AppErr)
}
