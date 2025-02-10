package repository

import (
	"context"

	"github.com/Marlliton/speisekarte/internal/entity/category"
	"github.com/Marlliton/speisekarte/pkg/apperr"
	"github.com/Marlliton/speisekarte/pkg/id"
)

type CategoryRepository interface {
	Save(ctx context.Context, category *category.Category) *apperr.AppErr
	FindByID(ctx context.Context, id id.ID) (*category.Category, *apperr.AppErr)
	FindAll(ctx context.Context) ([]*category.Category, *apperr.AppErr)
	Delete(ctx context.Context, id id.ID) *apperr.AppErr
	Update(ctx context.Context, id id.ID, category *category.Category) *apperr.AppErr
}
