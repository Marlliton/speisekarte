package repository

import (
	"context"

	"github.com/Marlliton/speisekarte/internal/entity/addon"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/Marlliton/speisekarte/pkg/id"
)

type AddOnRepository interface {
	Save(ctx context.Context, addon *addon.AddOn) *apperr.AppErr
	FindByID(ctx context.Context, id id.ID) (*addon.AddOn, *apperr.AppErr)
	FindAll(ctx context.Context) ([]*addon.AddOn, *apperr.AppErr)
	Delete(ctx context.Context, id id.ID) *apperr.AppErr
}
