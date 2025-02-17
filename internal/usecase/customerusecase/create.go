package customerusecase

import (
	"context"
	"time"

	"github.com/Marlliton/speisekarte/internal/entity/customer"
	"github.com/Marlliton/speisekarte/internal/repository"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/Marlliton/speisekarte/pkg/id"
)

type CustomerInput struct {
	Name  string
	Phone string
}

type CustomerOutput struct {
	ID        id.ID
	Name      string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type customerUseCase struct {
	repo repository.CustomerRepository
}

func New(repo repository.CustomerRepository) *customerUseCase {
	return &customerUseCase{repo: repo}
}

func (uc *customerUseCase) Execute(ctx context.Context, input CustomerInput) (*customer.Customer, *apperr.AppErr) {
	c, errs := customer.New(input.Name, input.Phone)

	if len(errs) > 0 {
		appErr := apperr.New("Invalid input").WithCode(apperr.INVALID_INPUT)
		for _, err := range errs {
			appErr.WithReason(err.Message, err.Field)
		}

		return nil, appErr
	}

	existingCustomer, err := uc.repo.FindByPhone(ctx, c.Phone)
	if err != nil {
		return nil, apperr.New("Failed to check customer existence").WithCode(apperr.INTERNAL)
	}
	if existingCustomer != nil {
		return nil, apperr.New("Customer already exists").WithCode(apperr.DUPLICATED)
	}

	if err := uc.repo.Create(ctx, c); err != nil {
		return nil, apperr.New("failed to create customer").WithCode(apperr.INTERNAL)
	}

	return c, nil
}
