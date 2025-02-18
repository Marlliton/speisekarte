package customer

import (
	"time"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/fail"
	"github.com/Marlliton/validator/rule"
)

type Customer struct {
	ID        id.ID
	Name      string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(name, phone string) (*Customer, []*fail.Error) {
	c := &Customer{
		ID:        id.New(),
		Name:      name,
		Phone:     phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if ok, errs := c.validate(); !ok {
		return nil, errs
	}

	return c, nil
}

func (c *Customer) validate() (bool, []*fail.Error) {
	v := validator.New()
	v.Add("ID", rule.Rules{rule.Required()})
	v.Add("Name", rule.Rules{rule.Required(), rule.MinLength(3), rule.MaxLength(100)})
	v.Add("Phone", rule.Rules{rule.Required(), rule.ValidPhoneNumber()})

	if errs := v.Validate(*c); len(errs) > 0 {
		return false, errs
	}

	return true, nil
}
