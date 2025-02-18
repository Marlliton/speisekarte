package customerusecase

import (
	"context"
	"testing"

	"github.com/Marlliton/speisekarte/internal/entity/customer"
	"github.com/Marlliton/speisekarte/internal/infrastructure/repository/inmemory"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/stretchr/testify/assert"
)

func TestCustomerUseCase_Execute(t *testing.T) {
	ctx := context.Background()
	phone := "+5511999999999"
	setup := func() *customerUseCase {
		repo := inmemory.NewInMemoryCustomerRepository()
		return New(repo)
	}

	t.Run("sould create and find by phone a new customer successfully", func(t *testing.T) {
		uc := setup()

		input := CustomerInput{
			Name:  "John Doe",
			Phone: phone,
		}

		result, err := uc.Execute(ctx, input)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, input.Name, result.Name)
		assert.Equal(t, input.Phone, result.Phone)
		assert.NotEmpty(t, result.UpdatedAt)
		assert.NotEmpty(t, result.CreatedAt)

		savedCustomer, err := uc.repo.FindByPhone(ctx, phone)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, input.Name, savedCustomer.Name)
		assert.Equal(t, input.Phone, savedCustomer.Phone)
	})

	t.Run("should retrun error when customer already exists", func(t *testing.T) {
		uc := setup()

		input := CustomerInput{
			Name:  "John Doe",
			Phone: phone,
		}

		existingCustomer, _ := customer.New("Existing Customer", phone)
		_ = uc.repo.Create(ctx, existingCustomer)

		result, err := uc.Execute(ctx, input)
		assert.NotNil(t, err)
		assert.Equal(t, CustomerOutput{}, result)
		assert.Equal(t, apperr.DUPLICATED, err.Code)
	})

	t.Run("should retorn error when input is invalid", func(t *testing.T) {
		uc := setup()

		input := CustomerInput{
			Name:  "",
			Phone: "",
		}

		_, err := uc.Execute(ctx, input)
		assert.NotNil(t, err)
		assert.Equal(t, apperr.INVALID_INPUT, err.Code)
	})
}
