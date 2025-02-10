package repository

import (
	"context"

	"github.com/Marlliton/speisekarte/internal/entity/menu"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/Marlliton/speisekarte/pkg/id"
)

type MenuRepository interface {
	Save(ctx context.Context, menu *menu.Menu) *apperr.AppErr
	FindByID(ctx context.Context, id id.ID) (*menu.Menu, *apperr.AppErr)
	FindAll(ctx context.Context) ([]*menu.Menu, *apperr.AppErr)
	Delete(ctx context.Context, id id.ID) *apperr.AppErr
	Update(ctx context.Context, id id.ID, menu *menu.Menu) *apperr.AppErr
	GetCategoriesByMenuID(ctx context.Context, menuID id.ID) ([]*menu.Menu, *apperr.AppErr)
}
