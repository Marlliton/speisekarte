package inmemory

import (
	"context"
	"testing"

	"github.com/Marlliton/speisekarte/internal/entity/addon"
	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestAddOnRepository(t *testing.T) {
	repo := NewInMemoryAddOnRepository()
	ctx := context.Background()

	t.Run("should save and find an addon", func(t *testing.T) {
		a := &addon.AddOn{ID: id.New(), Name: "Extra Cheese", Price: 200}
		err := repo.Create(ctx, a)
		assert.Nil(t, err)

		found, err := repo.FindByID(ctx, a.ID)
		assert.Nil(t, err)
		assert.Equal(t, a.ID, found.ID)
	})

	t.Run("should return error when addon is not found", func(t *testing.T) {
		_, err := repo.FindByID(ctx, id.New())
		assert.NotNil(t, err)
	})

	t.Run("should delete an addon", func(t *testing.T) {
		a := &addon.AddOn{ID: id.New(), Name: "Bacon", Price: 150}
		repo.Create(ctx, a)

		err := repo.Delete(ctx, a.ID)
		assert.Nil(t, err)

		_, err = repo.FindByID(ctx, a.ID)
		assert.NotNil(t, err)
	})
}
