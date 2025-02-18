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

func toCustomerOutput(c *customer.Customer) CustomerOutput {
	return CustomerOutput{
		ID:        c.ID,
		Name:      c.Name,
		Phone:     c.Phone,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func (uc *customerUseCase) Execute(ctx context.Context, input CustomerInput) (CustomerOutput, *apperr.AppErr) {
	c, errs := customer.New(input.Name, input.Phone)
	emptyOutput := CustomerOutput{}

	if len(errs) > 0 {
		appErr := apperr.New("Invalid input").WithCode(apperr.INVALID_INPUT)
		for _, err := range errs {
			appErr.WithReason(err.Message, err.Field)
		}

		return emptyOutput, appErr
	}

	existingCustomer, err := uc.repo.FindByPhone(ctx, c.Phone)
	if err != nil {
		return emptyOutput, apperr.New("Failed to check customer existence").WithCode(apperr.INTERNAL)
	}
	if existingCustomer != nil {
		return emptyOutput, apperr.New("Customer already exists").WithCode(apperr.DUPLICATED)
	}

	if err := uc.repo.Create(ctx, c); err != nil {
		return emptyOutput, apperr.New("failed to create customer").WithCode(apperr.INTERNAL)
	}

	return toCustomerOutput(c), nil
}
