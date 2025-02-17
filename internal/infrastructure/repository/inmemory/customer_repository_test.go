package inmemory

import (
	"context"
	"testing"

	"github.com/Marlliton/speisekarte/internal/entity/customer"
	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestCustomerRepository(t *testing.T) {
	repo := NewInMemoryCustomerRepository()
	ctx := context.Background()

	t.Run("should save and find a customer", func(t *testing.T) {
		c := &customer.Customer{
			ID:    id.New(),
			Name:  "John Doe",
			Phone: "+5511999999999",
		}

		err := repo.Create(ctx, c)
		assert.Nil(t, err)

		found, err := repo.FindByPhone(ctx, c.Phone)
		assert.Nil(t, err)
		assert.Equal(t, c.ID, found.ID)
		assert.Equal(t, c.Name, found.Name)
		assert.Equal(t, c.Phone, found.Phone)
	})

	t.Run("should return error when customer is not found", func(t *testing.T) {
		_, err := repo.FindByPhone(ctx, id.New().String())
		assert.NotNil(t, err)
		assert.Equal(t, "not found", err.Message)
	})
}
