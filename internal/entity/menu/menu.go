package menu

import (
	"time"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/fail"
	"github.com/Marlliton/validator/rule"
)

type Menu struct {
	ID         id.ID
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Categories []id.ID
}

func New(name string, categories ...id.ID) (*Menu, []*fail.Error) {
	m := &Menu{
		ID:         id.New(),
		Name:       name,
		CreatedAt:  time.Now(),
		Categories: categories,
	}

	ok, errs := m.validate()
	if !ok {
		return nil, errs
	}

	return m, nil
}

func (m *Menu) validate() (bool, []*fail.Error) {
	v := validator.New()
	v.Add("ID", rule.Rules{rule.Required()})
	v.Add("Name", rule.Rules{rule.Required(), rule.MinLength(3), rule.MaxLength(30)})

	errs := v.Validate(*m)
	if errs != nil {
		return false, errs
	}

	return true, nil
}
