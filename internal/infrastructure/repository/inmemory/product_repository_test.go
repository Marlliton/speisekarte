package inmemory

import (
	"context"
	"testing"

	"github.com/Marlliton/speisekarte/internal/entity/product"
	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestProductRepository(t *testing.T) {
	var repo *inMemoryProductRepository
	ctx := context.Background()

	setup := func() {
		repo = NewInMemoryProductRepository()
	}

	t.Run("should create a product", func(t *testing.T) {
		setup()
		product := &product.Product{
			ID:   id.New(),
			Name: "Test Product",
		}

		err := repo.Create(ctx, product)
		assert.Nil(t, err)

		foundProduct, err := repo.FindByID(ctx, product.ID)
		assert.Nil(t, err)
		assert.Equal(t, product.ID, foundProduct.ID)
		assert.Equal(t, product.Name, foundProduct.Name)
	})

	t.Run("should find a product by ID", func(t *testing.T) {
		setup()
		product := &product.Product{
			ID:   id.New(),
			Name: "Test Product",
		}

		repo.Create(ctx, product)

		foundProduct, err := repo.FindByID(ctx, product.ID)
		assert.Nil(t, err)
		assert.Equal(t, product.ID, foundProduct.ID)
		assert.Equal(t, product.Name, foundProduct.Name)
	})

	t.Run("should return error when product is not found", func(t *testing.T) {
		setup()
		nonExistentID := id.New()

		_, err := repo.FindByID(ctx, nonExistentID)
		assert.NotNil(t, err)
		assert.Equal(t, "message: not found", err.Error())
	})

	t.Run("should find all products", func(t *testing.T) {
		setup()
		product1 := &product.Product{ID: id.New(), Name: "Product 1"}
		product2 := &product.Product{ID: id.New(), Name: "Product 2"}

		repo.Create(ctx, product1)
		repo.Create(ctx, product2)

		products, err := repo.FindAll(ctx)
		assert.Nil(t, err)
		assert.Len(t, products, 2)
		assert.Equal(t, product1.Name, products[0].Name)
		assert.Equal(t, product2.Name, products[1].Name)
	})

	t.Run("should delete a product", func(t *testing.T) {
		setup()
		product := &product.Product{
			ID:   id.New(),
			Name: "Test Product",
		}

		repo.Create(ctx, product)

		err := repo.Delete(ctx, product.ID)
		assert.Nil(t, err)

		_, err = repo.FindByID(ctx, product.ID)
		assert.NotNil(t, err)
		assert.Equal(t, "message: not found", err.Error())
	})

	t.Run("should return error when deleting a non-existent product", func(t *testing.T) {
		setup()
		nonExistentID := id.New()

		err := repo.Delete(ctx, nonExistentID)
		assert.NotNil(t, err)
		assert.Equal(t, "message: not found", err.Error())
	})

	t.Run("should update a product", func(t *testing.T) {
		setup()
		prod := &product.Product{
			ID:   id.New(),
			Name: "Old Name",
		}

		repo.Create(ctx, prod)

		updatedProduct := &product.Product{
			ID:   prod.ID,
			Name: "New Name",
		}

		err := repo.Update(ctx, prod.ID, updatedProduct)
		assert.Nil(t, err)

		foundProduct, err := repo.FindByID(ctx, prod.ID)
		assert.Nil(t, err)
		assert.Equal(t, "New Name", foundProduct.Name)
	})

	t.Run("should return error when updating a non-existent product", func(t *testing.T) {
		setup()
		nonExistentID := id.New()
		updatedProduct := &product.Product{
			ID:   nonExistentID,
			Name: "New Name",
		}

		err := repo.Update(ctx, nonExistentID, updatedProduct)
		assert.NotNil(t, err)
	})

	t.Run("should return error when updating with nil product", func(t *testing.T) {
		setup()
		productID := id.New()

		err := repo.Update(ctx, productID, nil)
		assert.NotNil(t, err)
	})
}
