package customer

import (
	"testing"
	"time"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestCustomer_New(t *testing.T) {
	name := "John Doe"
	phone := "+5511999999999"

	t.Run("should create a customer", func(t *testing.T) {
		customer, errs := New(name, phone)

		assert.NotNil(t, customer)
		assert.Nil(t, errs)
		assert.NotEmpty(t, customer.ID)
		assert.Equal(t, name, customer.Name)
		assert.Equal(t, phone, customer.Phone)
	})

	t.Run("should not create a customer with empty name", func(t *testing.T) {
		customer, errs := New("", phone)

		assert.Nil(t, customer)
		assert.NotNil(t, errs)
		assert.Equal(t, 2, len(errs))
		assert.Equal(t, "Name", errs[0].Field)
	})

	t.Run("should not create a customer with name shorter than 3 characters", func(t *testing.T) {
		customer, errs := New("Jo", phone)

		assert.Nil(t, customer)
		assert.NotNil(t, errs)
		assert.Equal(t, 1, len(errs))
		assert.Equal(t, "Name", errs[0].Field)
	})

	t.Run("should not create a customer with name longer than 100 characters", func(t *testing.T) {
		longName := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris."
		customer, errs := New(longName, phone)

		assert.Nil(t, customer)
		assert.NotNil(t, errs)
		assert.Equal(t, 1, len(errs))
		assert.Equal(t, "Name", errs[0].Field)
	})

	t.Run("should not create a customer with empty phone", func(t *testing.T) {
		customer, errs := New(name, "")

		assert.Nil(t, customer)
		assert.NotNil(t, errs)
		assert.Equal(t, 2, len(errs))
		assert.Equal(t, "Phone", errs[0].Field)
	})

	t.Run("should not create a customer with invalid phone number", func(t *testing.T) {
		customer, errs := New(name, "123456")

		assert.Nil(t, customer)
		assert.NotNil(t, errs)
		assert.Equal(t, 1, len(errs))
		assert.Equal(t, "Phone", errs[0].Field)
	})
}

func TestCustomer_Validate(t *testing.T) {
	t.Run("should validate a customer with valid fields", func(t *testing.T) {
		customer := &Customer{
			ID:        id.New(),
			Name:      "John Doe",
			Phone:     "+5511999999999",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		ok, errs := customer.validate()

		assert.True(t, ok)
		assert.Nil(t, errs)
	})

	t.Run("should not validate a customer with empty ID", func(t *testing.T) {
		customer := &Customer{
			ID:        id.ID{},
			Name:      "John Doe",
			Phone:     "+5511999999999",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		ok, errs := customer.validate()

		assert.False(t, ok)
		assert.NotNil(t, errs)
		assert.Equal(t, 1, len(errs))
		assert.Equal(t, "ID", errs[0].Field)
	})

	t.Run("should not validate a customer with empty name", func(t *testing.T) {
		customer := &Customer{
			ID:        id.New(),
			Name:      "",
			Phone:     "+5511999999999",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		ok, errs := customer.validate()

		assert.False(t, ok)
		assert.NotNil(t, errs)
		assert.Equal(t, 2, len(errs))
		assert.Equal(t, "Name", errs[0].Field)
	})

	t.Run("should not validate a customer with invalid phone number", func(t *testing.T) {
		customer := &Customer{
			ID:        id.New(),
			Name:      "John Doe",
			Phone:     "123456",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		ok, errs := customer.validate()

		assert.False(t, ok)
		assert.NotNil(t, errs)
		assert.Equal(t, 1, len(errs))
		assert.Equal(t, "Phone", errs[0].Field)
	})
}
