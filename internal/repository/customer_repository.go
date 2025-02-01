package repository

import (
	"context"

	"github.com/Marlliton/speisekarte/internal/entity/customer"
	"github.com/Marlliton/speisekarte/pkg/id"
)

type CustomerRepository interface {
	Save(ctx context.Context, customer *customer.Customer) error
	FindByID(ctx context.Context, id id.ID) (*customer.Customer, error)
}
