package repository

import (
	"booking-api/internal/domain/table"
	"context"
)

type TableRepository interface{
	Create(ctx context.Context, t *table.Table) error
	GetByRestaurantID(ctx context.Context, restaurant_id int64) ([]table.Table, error)
	GetByID(ctx context.Context, id int64) (*table.Table, error)
}