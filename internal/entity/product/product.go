package product

import (
	"fmt"
	"time"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/fail"
	"github.com/Marlliton/validator/rule"
)

type Product struct {
	ID          id.ID
	Name        string
	Description string
	ImageURL    string
	Price       int
	Available   bool
	// AddOns      []id.ID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CategoryID id.ID
}

func New(
	name, description, imageURL string, price int, available bool, categoryID id.ID,
) (*Product, []*fail.Error) {
	p := &Product{
		ID:          id.New(),
		Name:        name,
		Description: description,
		ImageURL:    imageURL,
		Price:       price,
		Available:   available,
		CategoryID:  categoryID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	ok, errs := p.Validate()
	if !ok {
		return nil, errs
	}

	return p, nil
}

func (p *Product) DisplayPrice() string {
	return fmt.Sprintf("%.2f", float64(p.Price)/100)
}

func (p *Product) Validate() (bool, []*fail.Error) {
	v := validator.New()
	v.Add("ID", rule.Rules{rule.Required()})
	v.Add("Name", rule.Rules{
		rule.Required(),
		rule.MinLength(3),
		rule.MaxLength(50),
	})
	v.Add("Description", rule.Rules{
		rule.Required(),
		rule.MinLength(5),
		rule.MaxLength(300),
	})
	v.Add("ImageURL", rule.Rules{rule.Required(), rule.ValidURL()})
	v.Add("Price", rule.Rules{rule.Required(), rule.MinValue(0)})
	v.Add("Available", rule.Rules{rule.Bool()})
	v.Add("CategoryID", rule.Rules{rule.Required()})

	errs := v.Validate(*p)
	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}
