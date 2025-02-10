package menu

import (
	"time"

	"github.com/Marlliton/speisekarte/pkg/id"
	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/fail"
	"github.com/Marlliton/validator/rule"
)

// TODO: Adicionar comportamentos => AddCategory, RemoveCategory, GetCategories

type Menu struct {
	ID         id.ID
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	categories []id.ID
}

func New(name string, categories ...id.ID) (*Menu, []*fail.Error) {
	m := &Menu{
		ID:         id.New(),
		Name:       name,
		CreatedAt:  time.Now(),
		categories: categories,
	}

	ok, errs := m.validate()
	if !ok {
		return nil, errs
	}

	return m, nil
}

func (m *Menu) GetCategories() []id.ID {
	categories := make([]id.ID, len(m.categories))
	copy(categories, m.categories)

	return categories
}

func (m *Menu) AddCategory(categoryID id.ID) {
	for _, existingID := range m.categories {
		if existingID == categoryID {
			return
		}
	}

	m.categories = append(m.categories, categoryID)
	m.UpdatedAt = time.Now()
}

func (m *Menu) RemoveCategory(categoryID id.ID) {
	for i, ID := range m.categories {
		if ID == categoryID {
			m.categories = append(m.categories[:i], m.categories[i+1:]...)
		}
	}
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
