package inmemory

import (
	"context"
	"testing"

	"github.com/Marlliton/speisekarte/internal/entity/category"
	"github.com/Marlliton/speisekarte/internal/entity/menu"
	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestMenuRepository(t *testing.T) {
	var repo *MenuRepository
	var categoryRepo *CategoryRepository
	ctx := context.Background()

	setup := func() {
		categoryRepo = NewCategoryRepository()
		repo = NewMenuRepository(categoryRepo)
	}

	t.Run("should save and find a menu", func(t *testing.T) {
		setup()
		menu, errs := menu.New("Breakfast Menu")
		assert.Empty(t, errs)
		assert.NotNil(t, menu)

		err := repo.Create(ctx, menu)
		assert.Nil(t, err)

		found, err := repo.FindByID(ctx, menu.ID)
		assert.Nil(t, err)
		assert.Equal(t, menu.ID, found.ID)
		assert.Equal(t, menu.Name, found.Name)
	})

	t.Run("should return error when menu is not found", func(t *testing.T) {
		setup()
		_, err := repo.FindByID(ctx, id.New())
		assert.NotNil(t, err)
		assert.Equal(t, "message: not found", err.Error())
	})

	t.Run("should find all menus", func(t *testing.T) {
		setup()
		menu1, errs1 := menu.New("Lunch Menu")
		assert.Empty(t, errs1)
		menu2, errs2 := menu.New("Dinner Menu")
		assert.Empty(t, errs2)

		repo.Create(ctx, menu1)
		repo.Create(ctx, menu2)

		menus, err := repo.FindAll(ctx)
		assert.Nil(t, err)
		assert.Len(t, menus, 2)
	})

	t.Run("should delete a menu", func(t *testing.T) {
		setup()
		menu, errs := menu.New("Dessert Menu")
		assert.Empty(t, errs)

		repo.Create(ctx, menu)

		err := repo.Delete(ctx, menu.ID)
		assert.Nil(t, err)

		_, err = repo.FindByID(ctx, menu.ID)
		assert.NotNil(t, err)
		assert.Equal(t, "message: not found", err.Error())
	})

	t.Run("should update a menu", func(t *testing.T) {
		setup()
		m, errs := menu.New("Brunch Menu")
		assert.Empty(t, errs)

		repo.Create(ctx, m)

		updatedMenu, errs := menu.New("Updated Brunch Menu")
		assert.Empty(t, errs)
		updatedMenu.ID = m.ID

		err := repo.Update(ctx, updatedMenu)
		assert.Nil(t, err)

		found, err := repo.FindByID(ctx, m.ID)
		assert.Nil(t, err)
		assert.Equal(t, "Updated Brunch Menu", found.Name)
	})

	t.Run("should return error when updating a non-existent menu", func(t *testing.T) {
		setup()
		nonExistentMenu, errs := menu.New("Non-existent Menu")
		assert.Empty(t, errs)

		err := repo.Update(ctx, nonExistentMenu)
		assert.NotNil(t, err)
		assert.Equal(t, "message: not found", err.Error())
	})

	t.Run("should get categories by menu ID", func(t *testing.T) {
		setup()
		cat1 := &category.Category{ID: id.New(), Name: "Appetizers"}
		cat2 := &category.Category{ID: id.New(), Name: "Main Course"}
		categoryRepo.Create(ctx, cat1)
		categoryRepo.Create(ctx, cat2)

		menu, errs := menu.New("Dinner Menu", cat1.ID, cat2.ID)
		assert.Empty(t, errs)

		repo.Create(ctx, menu)

		categories, err := repo.GetCategoriesByMenuID(ctx, menu.ID)
		assert.Nil(t, err)
		assert.Len(t, categories, 2)
		assert.Equal(t, cat1.Name, categories[0].Name)
		assert.Equal(t, cat2.Name, categories[1].Name)
	})

	t.Run("should return error when menu has non-existent categories", func(t *testing.T) {
		setup()
		nonExistentCatID := id.New()
		menu, errs := menu.New("Invalid Menu", nonExistentCatID)
		assert.Empty(t, errs)

		repo.Create(ctx, menu)

		_, err := repo.GetCategoriesByMenuID(ctx, menu.ID)
		assert.NotNil(t, err)
		assert.Equal(t, "message: not found", err.Error())
	})
}
