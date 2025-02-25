package productusecase

import (
	"context"
	"testing"
	"time"

	"github.com/Marlliton/speisekarte/internal/entity/product"
	"github.com/Marlliton/speisekarte/internal/infrastructure/repository/inmemory"
	"github.com/Marlliton/speisekarte/internal/repository"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func setupTest() (*productUseCase, repository.ProductRepository) {
	repo := inmemory.NewInMemoryProductRepository()
	uc := New(repo)
	return uc, repo
}

func TestCreate(t *testing.T) {
	uc, repo := setupTest()
	ctx := context.Background()

	t.Run("Create valid product", func(t *testing.T) {
		input := ProductInput{
			Name:        "Test Product",
			Description: "Test Description",
			ImageURL:    "http://example.com/image.jpg",
			Price:       1000,
			Available:   true,
			CategoryID:  id.New(),
		}

		output, err := uc.Create(ctx, input)

		assert.Nil(t, err)
		assert.NotEmpty(t, output.ID)
		assert.Equal(t, input.Name, output.Name)
		assert.Equal(t, input.Description, output.Description)

		prod, err := repo.FindByID(ctx, output.ID)
		assert.Nil(t, err)
		assert.NotNil(t, prod)
	})

	t.Run("Create invalid product", func(t *testing.T) {
		input := ProductInput{
			Name:        "", // invalid name
			Description: "Test Description",
			ImageURL:    "http://example.com/image.jpg",
			Price:       1000,
			Available:   true,
			CategoryID:  id.New(),
		}

		_, err := uc.Create(ctx, input)
		assert.NotNil(t, err)
		assert.Equal(t, apperr.INVALID_INPUT, err.Code)
	})
}

func TestFindByID(t *testing.T) {
	uc, repo := setupTest()
	ctx := context.Background()

	prod := &product.Product{
		ID:          id.New(),
		Name:        "Test Product",
		Description: "Test Description",
		ImageURL:    "http://example.com/image.jpg",
		Price:       1000,
		Available:   true,
		CategoryID:  id.New(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	repo.Create(ctx, prod)

	t.Run("Find existing product", func(t *testing.T) {
		output, err := uc.FindByID(ctx, prod.ID)
		assert.Nil(t, err)
		assert.Equal(t, prod.ID, output.ID)
		assert.Equal(t, prod.Name, output.Name)
	})

	t.Run("Find non-existing product", func(t *testing.T) {
		nonExistingID := id.New()
		_, err := uc.FindByID(ctx, nonExistingID)
		t.Log(err)
		assert.NotNil(t, err)
		assert.Equal(t, apperr.NOT_FOUND, err.Code)
	})
}

func TestFindAll(t *testing.T) {
	uc, repo := setupTest()
	ctx := context.Background()

	prod1 := &product.Product{ID: id.New(), Name: "Product 1"}
	prod2 := &product.Product{ID: id.New(), Name: "Product 2"}
	repo.Create(ctx, prod1)
	repo.Create(ctx, prod2)

	t.Run("Find all products", func(t *testing.T) {
		output, err := uc.FindAll(ctx)
		t.Log("ERRR", err)
		assert.Nil(t, err)
		assert.Len(t, output, 2)
	})

	t.Run("Find all when no products exist", func(t *testing.T) {
		repo.Delete(ctx, prod1.ID)
		repo.Delete(ctx, prod2.ID)

		output, err := uc.FindAll(ctx)
		assert.Nil(t, err)
		assert.Empty(t, output)
		assert.Len(t, output, 0)
	})
}

func TestDelete(t *testing.T) {
	uc, repo := setupTest()
	ctx := context.Background()

	prod := &product.Product{ID: id.New(), Name: "Test Product"}
	repo.Create(ctx, prod)

	t.Run("Delete existing product", func(t *testing.T) {
		err := uc.Delete(ctx, prod.ID)
		assert.Nil(t, err)

		_, err = repo.FindByID(ctx, prod.ID)
		assert.NotNil(t, err)
		assert.Equal(t, apperr.NOT_FOUND, err.Code)
	})

	t.Run("Delete non-existing product", func(t *testing.T) {
		nonExistingID := id.New()
		err := uc.Delete(ctx, nonExistingID)
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	uc, repo := setupTest()
	ctx := context.Background()

	prod, errs := product.New("Old name", "Old Description", "http://example.com/old.png", 500, false, id.New())
	assert.Nil(t, errs)
	repo.Create(ctx, prod)

	t.Run("Update existing product", func(t *testing.T) {
		input := ProductInput{
			Name:        "New Name",
			Description: "New Description",
			ImageURL:    "http://example.com/new.jpg",
			Price:       1000,
			Available:   true,
			CategoryID:  id.New(),
		}

		err := uc.Update(ctx, prod.ID, input)
		assert.Nil(t, err)

		updatedProd, err := repo.FindByID(ctx, prod.ID)
		assert.Nil(t, err)
		assert.Equal(t, input.Name, updatedProd.Name)
		assert.Equal(t, input.Description, updatedProd.Description)
	})

	t.Run("Update non-existing product", func(t *testing.T) {
		nonExistingID := id.New()
		input := ProductInput{
			Name:        "New Name",
			Description: "New Description",
			ImageURL:    "http://example.com/new.jpg",
			Price:       1000,
			Available:   true,
			CategoryID:  id.New(),
		}

		err := uc.Update(ctx, nonExistingID, input)
		assert.NotNil(t, err)
		assert.Equal(t, apperr.NOT_FOUND, err.Code)
	})

	t.Run("Update with invalid input", func(t *testing.T) {
		input := ProductInput{
			Name:        "", // invalid name
			Description: "New Description",
			ImageURL:    "http://example.com/new.jpg",
			Price:       1000,
			Available:   true,
			CategoryID:  id.New(),
		}

		err := uc.Update(ctx, prod.ID, input)
		assert.NotNil(t, err)
		assert.Equal(t, apperr.INVALID_INPUT, err.Code)
	})
}
