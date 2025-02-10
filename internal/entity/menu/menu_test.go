package menu

import (
	"testing"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestMenu_New(t *testing.T) {
	name := "Menu Name"
	t.Run("should create a menu", func(t *testing.T) {
		m, errs := New(name)
		assert.Nil(t, errs)
		assert.Equal(t, name, m.Name)
		assert.Len(t, m.GetCategories(), 0)
	})

	t.Run("should create a menu with categories", func(t *testing.T) {
		m, errs := New(name, id.New(), id.New())
		assert.Nil(t, errs)
		assert.Equal(t, name, m.Name)
		assert.Len(t, m.GetCategories(), 2)
	})
}

func TestMenu_FailToCreate(t *testing.T) {
	t.Run("should fail to create", func(t *testing.T) {
		m, errs := New("")
		assert.NotNil(t, errs)
		assert.Nil(t, m)
	})
}

func TestMenu_AddCategory(t *testing.T) {
	t.Run("should add a category to the menu", func(t *testing.T) {
		m, _ := New("Menu Name")
		categoryID := id.New()

		m.AddCategory(categoryID)
		assert.Len(t, m.GetCategories(), 1)
		assert.Equal(t, categoryID, m.GetCategories()[0])
	})

	t.Run("should not add a duplicate category", func(t *testing.T) {
		m, _ := New("Menu Name")
		categoryID := id.New()

		m.AddCategory(categoryID)
		m.AddCategory(categoryID) // Tenta adicionar a mesma categoria novamente

		assert.Len(t, m.GetCategories(), 1) // A categoria não deve ser duplicada
	})
}

func TestMenu_RemoveCategory(t *testing.T) {
	t.Run("should remove a category from the menu", func(t *testing.T) {
		m, _ := New("Menu Name")
		categoryID := id.New()

		m.AddCategory(categoryID)
		assert.Len(t, m.GetCategories(), 1)

		m.RemoveCategory(categoryID)
		assert.Len(t, m.GetCategories(), 0)
	})

	t.Run("should do nothing if the category does not exist", func(t *testing.T) {
		m, _ := New("Menu Name")
		categoryID := id.New()

		m.RemoveCategory(categoryID) // Tenta remover uma categoria que não existe
		assert.Len(t, m.GetCategories(), 0)
	})
}

func TestMenu_GetCategories(t *testing.T) {
	t.Run("should return a copy of the categories", func(t *testing.T) {
		m, _ := New("Menu Name")
		categoryID1 := id.New()
		categoryID2 := id.New()

		m.AddCategory(categoryID1)
		m.AddCategory(categoryID2)

		categories := m.GetCategories()
		assert.Len(t, categories, 2)
		assert.Equal(t, categoryID1, categories[0])
		assert.Equal(t, categoryID2, categories[1])

		// Modificar a cópia não deve afetar o menu original
		categories[0] = id.New()
		assert.NotEqual(t, categories[0], m.GetCategories()[0])
	})
}
