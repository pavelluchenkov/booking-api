package repository

import (
	"booking-api/internal/domain/restaurant"
	"context"
)

type RestaurantRepository interface {
	Create(ctx context.Context, r *restaurant.Restaurant) error
	GetAll(ctx context.Context) ([]restaurant.Restaurant, error)
	GetByID(ctx context.Context, id int64) (*restaurant.Restaurant, error)
}
