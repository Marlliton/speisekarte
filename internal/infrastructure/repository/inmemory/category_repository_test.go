package inmemory

import (
	"context"
	"testing"

	"github.com/Marlliton/speisekarte/internal/entity/category"
	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestCategoryRepository(t *testing.T) {
	var repo *CategoryRepository
	ctx := context.Background()

	setup := func() {
		repo = NewCategoryRepository()
	}

	t.Run("should save and find a category", func(t *testing.T) {
		setup()
		cat := &category.Category{ID: id.New(), Name: "Drinks"}
		err := repo.Save(ctx, cat)
		assert.Nil(t, err)

		found, err := repo.FindByID(ctx, cat.ID)
		assert.Nil(t, err)
		assert.Equal(t, cat.ID, found.ID)
		assert.Equal(t, cat.Name, found.Name)
	})

	t.Run("should return error when category is not found", func(t *testing.T) {
		setup()
		_, err := repo.FindByID(ctx, id.New())
		assert.NotNil(t, err)
	})

	t.Run("should find all categories", func(t *testing.T) {
		setup()
		cat1 := &category.Category{ID: id.New(), Name: "Appetizers"}
		cat2 := &category.Category{ID: id.New(), Name: "Main Course"}
		repo.Save(ctx, cat1)
		repo.Save(ctx, cat2)

		categories, err := repo.FindAll(ctx)
		t.Log(categories)
		assert.Nil(t, err)
		assert.Len(t, categories, 2)
	})

	t.Run("should delete a category", func(t *testing.T) {
		setup()
		cat := &category.Category{ID: id.New(), Name: "Desserts"}
		repo.Save(ctx, cat)

		err := repo.Delete(ctx, cat.ID)
		assert.Nil(t, err)

		_, err = repo.FindByID(ctx, cat.ID)
		assert.NotNil(t, err)
	})

	t.Run("should update a category", func(t *testing.T) {
		setup()
		cat := &category.Category{ID: id.New(), Name: "Beverages"}
		repo.Save(ctx, cat)

		updateFields := map[string]interface{}{
			"Name": "Hot Beverages",
		}
		err := repo.Update(ctx, cat.ID, updateFields)
		assert.Nil(t, err)

		updatedCat, err := repo.FindByID(ctx, cat.ID)
		assert.Nil(t, err)
		assert.Equal(t, "Hot Beverages", updatedCat.Name)
	})

	t.Run("should return error when updating a non-existent category", func(t *testing.T) {
		setup()
		updateFields := map[string]interface{}{
			"Name": "Non-existent",
		}
		err := repo.Update(ctx, id.New(), updateFields)
		assert.NotNil(t, err)
	})

	t.Run("should return error when updating with invalid field", func(t *testing.T) {
		setup()
		cat := &category.Category{ID: id.New(), Name: "Snacks"}
		repo.Save(ctx, cat)

		updateFields := map[string]interface{}{
			"InvalidField": "Some value",
		}
		err := repo.Update(ctx, cat.ID, updateFields)
		assert.NotNil(t, err)
	})
}
