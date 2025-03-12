package orderusecase

import (
	"context"
	"testing"

	"github.com/Marlliton/speisekarte/internal/infrastructure/repository/inmemory"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestOrderUseCase_Create_Success(t *testing.T) {
	repo := inmemory.NewInMemoryOrderRepository()
	uc := New(repo)

	customerID := id.New()
	input := OrderInput{
		CustomerID: customerID,
		Items: []struct {
			ProductID id.ID
			Quantity  int
			Price     int
			AddOns    []struct {
				Name     string
				Price    int
				Quantity int
			}
		}{
			{
				ProductID: id.New(),
				Quantity:  2,
				Price:     1000,
				AddOns: []struct {
					Name     string
					Price    int
					Quantity int
				}{
					{
						Name:     "Extra Cheese",
						Price:    200,
						Quantity: 1,
					},
				},
			},
		},
	}

	t.Run("should create an order successfully", func(t *testing.T) {
		err := uc.Create(context.Background(), input)

		assert.Nil(t, err)

		orders, err := repo.FindAllByCustomerID(customerID)
		assert.Nil(t, err)
		assert.Len(t, orders, 1)
		assert.Equal(t, customerID, orders[0].CustomerID)
		assert.Len(t, orders[0].Items, 1)
		assert.Equal(t, 2, orders[0].Items[0].Quantity)
		assert.Equal(t, 1000, orders[0].Items[0].Price)
		assert.Len(t, orders[0].Items[0].AddOns, 1)
		assert.Equal(t, "Extra Cheese", orders[0].Items[0].AddOns[0].Name)
	})
}

func TestOrderUseCase_Create_Fail(t *testing.T) {
	repo := inmemory.NewInMemoryOrderRepository()
	uc := New(repo)

	t.Run("should fail to create an order with invalid customer id", func(t *testing.T) {
		input := OrderInput{
			CustomerID: id.ID{}, // Invalid customer ID
			Items: []struct {
				ProductID id.ID
				Quantity  int
				Price     int
				AddOns    []struct {
					Name     string
					Price    int
					Quantity int
				}
			}{
				{
					ProductID: id.New(),
					Quantity:  2,
					Price:     1000,
				},
			},
		}

		err := uc.Create(context.Background(), input)

		assert.NotNil(t, err)
		assert.Equal(t, apperr.INVALID_INPUT, err.Code)
	})

	t.Run("should fail to create an order with invalid product id", func(t *testing.T) {
		input := OrderInput{
			CustomerID: id.New(),
			Items: []struct {
				ProductID id.ID
				Quantity  int
				Price     int
				AddOns    []struct {
					Name     string
					Price    int
					Quantity int
				}
			}{
				{
					ProductID: id.ID{}, // Invalid product ID
					Quantity:  2,
					Price:     1000,
				},
			},
		}

		err := uc.Create(context.Background(), input)

		assert.NotNil(t, err)
		assert.Equal(t, apperr.INVALID_INPUT, err.Code)
	})

	t.Run("should fail to create an order with invalid add-on", func(t *testing.T) {
		input := OrderInput{
			CustomerID: id.New(),
			Items: []struct {
				ProductID id.ID
				Quantity  int
				Price     int
				AddOns    []struct {
					Name     string
					Price    int
					Quantity int
				}
			}{
				{
					ProductID: id.New(),
					Quantity:  2,
					Price:     1000,
					AddOns: []struct {
						Name     string
						Price    int
						Quantity int
					}{
						{
							Name:     "", // Invalid add-on name
							Price:    200,
							Quantity: 1,
						},
					},
				},
			},
		}

		err := uc.Create(context.Background(), input)

		assert.NotNil(t, err)
		assert.Equal(t, apperr.INVALID_INPUT, err.Code)
	})
}
