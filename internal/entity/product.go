package entity

import (
	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/rule"
	"github.com/Marlliton/validator/validator_error"
)

type Product struct {
	ID          id.ID
	Name        string
	Description string
	ImageURL    string
	Price       int
	Available   bool
	CategoryID  id.ID
}

type Numeric interface{ int | float64 }

func NewProduct[T Numeric](
	name, description, imageURL string, price T, available bool, categoryID id.ID,
) (*Product, []*validator_error.ValidatorError) {
	priceInCents := convertToCents(price)
	p := &Product{
		ID:          id.New(),
		Name:        name,
		Description: description,
		ImageURL:    imageURL,
		Price:       priceInCents,
		Available:   available,
		CategoryID:  categoryID,
	}

	ok, errs := p.validate()
	if !ok {
		return nil, errs
	}

	return p, nil
}

func convertToCents[T Numeric](value T) int {
	switch v := any(value).(type) {
	case int:
		return v * 100
	case float64:
		return int(v) * 100
	default:
		return -1
	}
}

func (p *Product) validate() (bool, []*validator_error.ValidatorError) {
	v := validator.New()
	v.Add("ID", rule.Rules{
		rule.Required(),
	})
	v.Add("Name", rule.Rules{
		rule.Required(),
		rule.MinLength(3),
		rule.MaxLength(50),
	})
	v.Add("Description", rule.Rules{
		rule.Required(),
		rule.MinLength(15),
		rule.MaxLength(300),
	})
	v.Add("ImageURL", rule.Rules{
		rule.Required(),
		rule.ValidURL(),
	})
	v.Add("Price", rule.Rules{
		rule.Required(),
		rule.MinValue(0),
	})
	v.Add("Available", rule.Rules{
		rule.Bool(),
	})
	v.Add("CategoryID", rule.Rules{rule.Required()})

	errs := v.Validate(*p)
	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}
