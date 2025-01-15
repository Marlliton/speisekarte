package category

import (
	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/rule"
	"github.com/Marlliton/validator/validator_error"
)

type Category struct {
	ID   id.ID
	Name string
}

func NewCategory(name string) (*Category, []*validator_error.ValidatorError) {
	c := &Category{ID: id.New(), Name: name}

	ok, errs := c.validate()
	if !ok {
		return nil, errs
	}

	return c, nil
}

func (c *Category) validate() (bool, []*validator_error.ValidatorError) {
	v := validator.New()
	v.Add("ID", rule.Rules{rule.Required()})
	v.Add("Name", rule.Rules{rule.Required(), rule.MinLength(3), rule.MaxLength(30)})

	errs := v.Validate(*c)
	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}
