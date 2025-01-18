package cart

import (
	"testing"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestNewCartItem_Success(t *testing.T) {
	priceInCents := 900 // 9 reais
	c, errs := NewCartItem(id.New(), id.New(), priceInCents, 2)
	assert.Nil(t, errs)
	assert.NotNil(t, c)
	assert.Equal(t, 2, c.Quantity)
}

func TestNewCartItem_ValidationFails(t *testing.T) {
	// Missing cart ID and prod ID
	cartID := id.ID{}
	prodID := id.ID{}
	ci, errs := NewCartItem(cartID, prodID, 900, 0)
	assert.Nil(t, ci)
	assert.NotNil(t, errs)
}
