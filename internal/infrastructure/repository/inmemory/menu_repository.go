package inmemory

import (
	"context"
	"reflect"
	"sync"

	"github.com/Marlliton/speisekarte/internal/entity/category"
	"github.com/Marlliton/speisekarte/internal/entity/menu"
	"github.com/Marlliton/speisekarte/internal/repository"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/Marlliton/speisekarte/pkg/id"
)

type MenuRepository struct {
	sync.RWMutex
	menus        map[id.ID]*menu.Menu
	categoryRepo repository.CategoryRepository
}

func NewMenuRepository(categoryRepo repository.CategoryRepository) *MenuRepository {
	return &MenuRepository{menus: map[id.ID]*menu.Menu{}, categoryRepo: categoryRepo}
}

func (r *MenuRepository) Save(ctx context.Context, menu *menu.Menu) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	r.menus[menu.ID] = menu
	return nil
}

func (r *MenuRepository) FindByID(ctx context.Context, id id.ID) (*menu.Menu, *apperr.AppErr) {
	r.RLock()
	defer r.RUnlock()

	if _, exists := r.menus[id]; !exists {
		return nil, apperr.New("not found")
	}

	return r.menus[id], nil
}

func (r *MenuRepository) FindAll(ctx context.Context) ([]*menu.Menu, *apperr.AppErr) {
	r.RLock()
	defer r.RUnlock()

	menus := make([]*menu.Menu, 0, len(r.menus))
	for _, c := range r.menus {
		menus = append(menus, c)
	}

	return menus, nil
}

func (r *MenuRepository) Delete(ctx context.Context, id id.ID) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	if _, exists := r.menus[id]; !exists {
		return apperr.New("not found")
	}

	delete(r.menus, id)
	return nil
}

func (r *MenuRepository) Update(ctx context.Context, id id.ID, fields map[string]interface{}) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	cat, exists := r.menus[id]
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
func (r *MenuRepository) GetCategoriesByMenuID(ctx context.Context, menuID id.ID) ([]*category.Category, *apperr.AppErr) {
	r.RLock()
	defer r.RUnlock()

	categories := make([]*category.Category, 0)
	menu, err := r.FindByID(ctx, menuID)
	if err != nil {
		return nil, apperr.New("not found")
	}
	for _, catID := range menu.GetCategories() {
		cat, err := r.categoryRepo.FindByID(ctx, catID)
		if err != nil {
			return nil, apperr.New("not found")
		}
		categories = append(categories, cat)
	}

	return categories, nil
}
