package inmemory

import (
	"context"
	"sync"

	"github.com/Marlliton/speisekarte/internal/entity/addon"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/Marlliton/speisekarte/pkg/id"
)

type inMemoryAddOnRepository struct {
	sync.RWMutex
	addOns map[id.ID]*addon.AddOn
}

func NewInMemoryAddOnRepository() *inMemoryAddOnRepository {
	return &inMemoryAddOnRepository{addOns: make(map[id.ID]*addon.AddOn)}
}

func (r *inMemoryAddOnRepository) Create(ctx context.Context, a *addon.AddOn) *apperr.AppErr {
	r.Lock()
	defer r.Unlock()

	r.addOns[a.ID] = a
	return nil
}

func (r *inMemoryAddOnRepository) FindByID(ctx context.Context, id id.ID) (*addon.AddOn, *apperr.AppErr) {
	r.RLock()
	defer r.RUnlock()

	if _, exists := r.addOns[id]; !exists {
		return nil, apperr.New("not found")
	}

	return r.addOns[id], nil
}

func (r *inMemoryAddOnRepository) FindAll(ctx context.Context) ([]*addon.AddOn, *apperr.AppErr) {
	r.RLock()
	defer r.RUnlock()

	addons := make([]*addon.AddOn, 0, len(r.addOns))
	for _, addon := range r.addOns {
		addons = append(addons, addon)
	}

	return addons, nil
}

func (r *inMemoryAddOnRepository) Delete(ctx context.Context, id id.ID) *apperr.AppErr {

	r.RLock()
	defer r.RUnlock()

	if _, exists := r.addOns[id]; !exists {
		return apperr.New("not found")
	}

	delete(r.addOns, id)

	return nil
}
