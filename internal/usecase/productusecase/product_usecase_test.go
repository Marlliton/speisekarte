package productusecase

import (
	"context"
	"testing"

	"github.com/Marlliton/speisekarte/internal/infrastructure/repository/inmemory"
	"github.com/Marlliton/speisekarte/internal/repository"
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
}
