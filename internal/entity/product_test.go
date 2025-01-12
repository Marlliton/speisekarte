package entity

import (
	"testing"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestProduct_New(t *testing.T) {
	t.Run("should create a new product", func(t *testing.T) {
		p, errs := New("Product 1", "Product Description", "http://test.com", 19, false, id.New())
		t.Log(errs)
		assert.Nil(t, errs)
		assert.NotNil(t, p.Name)
		assert.NotNil(t, p.Description)
		assert.NotNil(t, p.ImageURL)
		assert.NotNil(t, p.Price)
		assert.False(t, p.Available)
	})
}
