package menu

import (
	"testing"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestMenu_New(t *testing.T) {
	name := "Menu Name"
	t.Run("should to create a menu", func(t *testing.T) {
		m, errs := New(name)
		assert.Nil(t, errs)
		assert.Equal(t, name, m.Name)
		assert.Len(t, m.Categories, 0)
	})

	t.Run("should to create a menu with categories", func(t *testing.T) {
		m, errs := New(name, id.New(), id.New())
		assert.Nil(t, errs)
		assert.Equal(t, name, m.Name)
		assert.Len(t, m.Categories, 2)
	})
}

func TestMenu_FailToCreate(t *testing.T) {
	t.Run("should fail to create", func(t *testing.T) {
		m, errs := New("")
		assert.NotNil(t, errs)
		assert.Nil(t, m)
	})
}
