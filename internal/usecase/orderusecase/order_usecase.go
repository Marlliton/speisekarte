package orderusecase

import (
	"context"

	"github.com/Marlliton/speisekarte/internal/entity/order"
	"github.com/Marlliton/speisekarte/internal/repository"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/Marlliton/speisekarte/pkg/id"
)

type OrderInput struct {
	CustomerID id.ID
	Items      []struct {
		ProductID id.ID
		Quantity  int
		Price     int
		AddOns    []struct {
			Name     string
			Price    int
			Quantity int
		}
	}
}

// type OrderOutput struct {
// 	Status string
// 	Total  int
//
// 	Items []struct {
// 		Product struct {
// 			Name        string
// 			Description string
// 			ImageUrl    string
// 		}
// 		AddOns []struct {
// 			OrderItemID string
// 			Name        string
// 			Price       int
// 			Quantity    int
// 		}
// 		Quantity int
// 		Price    int
// 	}
// }

type orderUseCase struct {
	repo repository.OrderRepository
}

func New(repo repository.OrderRepository) *orderUseCase {
	return &orderUseCase{repo: repo}
}

func (uc *orderUseCase) Create(ctx context.Context, input OrderInput) *apperr.AppErr {
	orderID := id.New()
	items := make([]*order.OrderItem, 0, len(input.Items))
	inputErrs := apperr.New("invalid-input").WithCode(apperr.INVALID_INPUT)

	for _, it := range input.Items {
		itemID := id.New()
		addons := make([]*order.AddOn, 0, len(it.AddOns))

		for _, ad := range it.AddOns {
			addon, errs := order.NewAddOn(itemID, ad.Name, ad.Price, ad.Quantity)
			if len(errs) > 0 {
				for _, err := range errs {
					inputErrs.WithReason(err.Message, err.Field)
				}
				return inputErrs
			}

			addons = append(addons, addon)
		}

		item, errs := order.OrderNewItem(itemID, orderID, it.ProductID, it.Quantity, it.Price, addons...)
		if len(errs) > 0 {
			for _, err := range errs {
				inputErrs.WithReason(err.Message, err.Field)
			}
			return inputErrs
		}
		items = append(items, item)
	}

	order, errs := order.New(orderID, input.CustomerID, items...)
	if len(errs) > 0 {
		for _, err := range errs {
			inputErrs.WithReason(err.Message, err.Field)
		}
		return inputErrs
	}

	err := uc.repo.Create(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

// func (uc *orderUseCase) FindAllByCustomerID(customerID id.ID) ([]*order.Order, *apperr.AppErr) {}
