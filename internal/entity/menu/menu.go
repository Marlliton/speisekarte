package menu

import (
	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/rule"
	"github.com/Marlliton/validator/validator_error"
)

type Menu struct {
	ID         id.ID
	Name       string
	Categories []id.ID
}

func New(name string, categories ...id.ID) (*Menu, []*validator_error.ValidatorError) {
	m := &Menu{
		ID:         id.New(),
		Name:       name,
		Categories: categories,
	}

	ok, errs := m.validate()
	if !ok {
		return nil, errs
	}

	return m, nil
}

func (m *Menu) validate() (bool, []*validator_error.ValidatorError) {
	v := validator.New()
	v.Add("ID", rule.Rules{rule.Required()})
	v.Add("Name", rule.Rules{rule.Required(), rule.MinLength(3), rule.MaxLength(30)})

	errs := v.Validate(*m)
	if errs != nil {
		return false, errs
	}

	return true, nil
}
