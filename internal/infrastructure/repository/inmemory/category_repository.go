package inmemory

import (
	"context"
	"reflect"
	"sync"

	"github.com/Marlliton/speisekarte/internal/entity/category"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/Marlliton/speisekarte/pkg/id"
)

type CategoryRepository struct {
	sync.RWMutex
	categories map[id.ID]*category.Category
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{categories: map[id.ID]*category.Category{}}
}

func (r *CategoryRepository) Save(ctx context.Context, category *category.Category) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	r.categories[category.ID] = category
	return nil
}

func (r *CategoryRepository) FindByID(ctx context.Context, id id.ID) (*category.Category, *apperr.AppErr) {
	r.RLock()
	defer r.RUnlock()

	if _, exists := r.categories[id]; !exists {
		return nil, apperr.New("not found")
	}

	return r.categories[id], nil
}

func (r *CategoryRepository) FindAll(ctx context.Context) ([]*category.Category, *apperr.AppErr) {
	r.RLock()
	defer r.RUnlock()

	categories := make([]*category.Category, 0, len(r.categories))
	for _, c := range r.categories {
		categories = append(categories, c)
	}

	return categories, nil
}

func (r *CategoryRepository) Delete(ctx context.Context, id id.ID) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.categories[id]; !exists {
		return apperr.New("not found")
	}

	delete(r.categories, id)
	return nil
}

func (r *CategoryRepository) Update(ctx context.Context, id id.ID, fields map[string]interface{}) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	cat, exists := r.categories[id]
	if !exists {
		return apperr.New("not found")
	}

	catValue := reflect.ValueOf(cat).Elem()
	for field, value := range fields {
		fieldValue := catValue.FieldByName(field)
		if !fieldValue.IsValid() {
			return apperr.New("invalid field")
		}
		if !fieldValue.CanSet() {
			return apperr.New("invalid field")
		}

		fieldValue.Set(reflect.ValueOf(value))
	}

	return nil
}
