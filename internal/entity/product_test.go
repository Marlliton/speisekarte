package entity

import (
	"testing"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestProduct_New(t *testing.T) {
	t.Run("should create a valid product", func(t *testing.T) {
		categoryID := id.New()
		p, errs := NewProduct("Product Name", "Product Description", "http://test.com", 19.65, true, categoryID)
		assert.Nil(t, errs)
		assert.NotNil(t, p.Name)
		assert.NotNil(t, p.Description)
		assert.NotNil(t, p.ImageURL)
		assert.NotNil(t, p.Price)
		assert.True(t, p.Available)
		assert.Equal(t, 1965, p.Price)
	})

	t.Run("should return display price", func(t *testing.T) {
		categoryID := id.New()
		p, _ := NewProduct("Product Name", "Product Description", "http://test.com", 19.65, true, categoryID)
		assert.Equal(t, 1965, p.Price)

		displayPrice := p.DisplayPrice()
		assert.Equal(t, "19.65", displayPrice)
	})

	tests := []struct {
		name           string
		description    string
		imageURL       string
		price          float64
		available      bool
		categoryID     id.ID
		expectedErrLen int
		label          string
	}{
		{
			name:           "",
			description:    "Valid Description",
			imageURL:       "http://test.com",
			price:          10.5,
			available:      true,
			categoryID:     id.New(),
			expectedErrLen: 2,
			label:          "should fail when Name is empty or too short",
		},
		{
			name:           "valid name",
			description:    "",
			imageURL:       "http://test.com",
			price:          10.5,
			available:      true,
			categoryID:     id.New(),
			expectedErrLen: 2,
			label:          "should fail when Description is empty or too short",
		},
		{
			name:           "valid name",
			description:    "Valid description",
			imageURL:       "",
			price:          10.5,
			available:      true,
			categoryID:     id.New(),
			expectedErrLen: 2,
			label:          "should fail when ImageURL is empty",
		},
		{
			name:           "valid name",
			description:    "Valid description",
			imageURL:       "ivalid_url",
			price:          10.5,
			available:      true,
			categoryID:     id.New(),
			expectedErrLen: 1,
			label:          "should fail when ImageURL is invalid",
		},
		{
			name:           "valid name",
			description:    "Valid description",
			imageURL:       "https://valid.com",
			price:          10.5,
			available:      true,
			categoryID:     id.ID{},
			expectedErrLen: 1,
			label:          "should fail when categoryID is invalid",
		},
	}

	for _, tc := range tests {
		t.Run(tc.label, func(t *testing.T) {
			_, errs := NewProduct(tc.name, tc.description, tc.imageURL, tc.price, tc.available, tc.categoryID)
			assert.NotNil(t, errs)
			assert.Len(t, errs, tc.expectedErrLen)
		})
	}
}
