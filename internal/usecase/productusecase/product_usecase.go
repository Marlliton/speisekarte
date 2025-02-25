package productusecase

import (
	"context"
	"fmt"
	"time"

	"github.com/Marlliton/speisekarte/internal/entity/product"
	"github.com/Marlliton/speisekarte/internal/repository"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/Marlliton/speisekarte/pkg/id"
)

type ProductInput struct {
	Name        string
	Description string
	ImageURL    string
	Price       int
	Available   bool
	CategoryID  id.ID
}
type ProductOutput struct {
	ID          id.ID
	Name        string
	Description string
	ImageURL    string
	Price       int
	Available   bool
	CategoryID  id.ID

	CreatedAt time.Time
	UpdatedAt time.Time
}

type productUseCase struct {
	repo repository.ProductRepository
}

func New(repo repository.ProductRepository) *productUseCase {
	return &productUseCase{repo}
}

func toProductOutput(p *product.Product) ProductOutput {
	return ProductOutput{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		ImageURL:    p.ImageURL,
		Price:       p.Price,
		Available:   p.Available,
		CategoryID:  p.CategoryID,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func (uc *productUseCase) Create(ctx context.Context, input ProductInput) (ProductOutput, *apperr.AppErr) {
	prod, errs := product.New(input.Name, input.Description, input.ImageURL, input.Price, input.Available, input.CategoryID)
	emptyOut := ProductOutput{}
	if len(errs) > 0 {
		appErr := apperr.New("invalid product input").WithCode(apperr.INVALID_INPUT)
		for _, e := range errs {
			appErr.WithReason(e.Message, e.Field)
		}

		return emptyOut, appErr
	}

	err := uc.repo.Create(ctx, prod)
	if err != nil {
		return emptyOut, apperr.New("failed to create product").WithCode(apperr.INTERNAL)
	}

	return toProductOutput(prod), nil
}

func (uc *productUseCase) FindByID(ctx context.Context, id id.ID) (ProductOutput, *apperr.AppErr) {
	prod, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		msg := fmt.Sprintf("failed to find product by id: %s", err.Error())
		return ProductOutput{}, apperr.New(msg).WithCode(apperr.INTERNAL)
	}

	if prod == nil {
		return ProductOutput{}, apperr.New("product not found").WithCode(apperr.NOT_FOUND)
	}

	return toProductOutput(prod), nil
}

func (uc *productUseCase) FindAll(ctx context.Context) ([]ProductOutput, *apperr.AppErr) {
	result, err := uc.repo.FindAll(ctx)
	if err != nil {
		msg := fmt.Sprintf("failed to get all product: %s", err.Error())
		return nil, apperr.New(msg).WithCode(apperr.INTERNAL)
	}

	output := make([]ProductOutput, 0, len(result))

	for i, p := range result {
		output[i] = toProductOutput(p)
	}

	return output, nil
}

func (uc *productUseCase) Delete(ctx context.Context, id id.ID) *apperr.AppErr {
	err := uc.repo.Delete(ctx, id)
	if err != nil {
		msg := fmt.Sprintf("failed to delete product: %s", err.Error())
		return apperr.New(msg).WithCode(apperr.INTERNAL)
	}

	return nil
}

func (uc *productUseCase) Update(ctx context.Context, id id.ID, input ProductInput) *apperr.AppErr {
	existingProd, err := uc.FindByID(ctx, id)
	if err != nil {
		return err
	}

	updatedProd := &product.Product{
		ID:          existingProd.ID,
		Name:        input.Name,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		Price:       input.Price,
		Available:   input.Available,
		CategoryID:  input.CategoryID,
		CreatedAt:   existingProd.CreatedAt,
		UpdatedAt:   time.Now(),
	}

	ok, errs := updatedProd.Validate()

	if !ok {
		appErr := apperr.New("invalid product input").WithCode(apperr.INVALID_INPUT)
		for _, e := range errs {
			appErr.WithReason(e.Message, e.Field)
		}

		return appErr
	}

	err = uc.repo.Update(ctx, id, updatedProd)
	if err != nil {
		msg := fmt.Sprintf("failed to update product: %s", err.Error())
		return apperr.New(msg).WithCode(apperr.INTERNAL)
	}

	return nil
}
