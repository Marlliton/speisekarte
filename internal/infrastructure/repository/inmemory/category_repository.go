package inmemory

import (
	"context"
	"fmt"
	"sync"

	"github.com/Marlliton/speisekarte/internal/entity/category"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/Marlliton/speisekarte/pkg/id"
)

type inMemoryCategoryRepository struct {
	sync.RWMutex
	categories map[id.ID]*category.Category
}

func NewInMemoryCategoryRepository() *inMemoryCategoryRepository {
	return &inMemoryCategoryRepository{categories: map[id.ID]*category.Category{}}
}

func (r *inMemoryCategoryRepository) Create(ctx context.Context, category *category.Category) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	r.categories[category.ID] = category
	return nil
}

func (r *inMemoryCategoryRepository) FindByID(ctx context.Context, id id.ID) (*category.Category, *apperr.AppErr) {
	r.RLock()
	defer r.RUnlock()

	if _, exists := r.categories[id]; !exists {
		return nil, apperr.New("not found")
	}

	return r.categories[id], nil
}

func (r *inMemoryCategoryRepository) FindAll(ctx context.Context) ([]*category.Category, *apperr.AppErr) {
	r.RLock()
	defer r.RUnlock()

	categories := make([]*category.Category, 0, len(r.categories))
	for _, c := range r.categories {
		categories = append(categories, c)
	}

	return categories, nil
}

func (r *inMemoryCategoryRepository) Delete(ctx context.Context, id id.ID) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.categories[id]; !exists {
		return apperr.New("not found")
	}

	delete(r.categories, id)
	return nil
}

func (r *inMemoryCategoryRepository) Update(ctx context.Context, id id.ID, category *category.Category) *apperr.AppErr {
	if category == nil {
		return apperr.New("category cannot be nil")
	}

	r.Lock()
	defer r.Unlock()

	existingCategory, exists := r.categories[id]
	if !exists {
		return apperr.New(fmt.Sprintf("category with ID %s not found", id))
	}

	existingCategory.Name = category.Name

	return nil
}
