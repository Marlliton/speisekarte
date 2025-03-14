package category

import (
	"time"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/fail"
	"github.com/Marlliton/validator/rule"
)

type Category struct {
	ID        id.ID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(name string) (*Category, []*fail.Error) {
	c := &Category{ID: id.New(), Name: name, CreatedAt: time.Now()}

	ok, errs := c.validate()
	if !ok {
		return nil, errs
	}

	return c, nil
}

func (c *Category) validate() (bool, []*fail.Error) {
	v := validator.New()
	v.Add("ID", rule.Rules{rule.Required()})
	v.Add("Name", rule.Rules{rule.Required(), rule.MinLength(3), rule.MaxLength(30)})

	errs := v.Validate(*c)
	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}
