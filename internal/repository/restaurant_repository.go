package repository

import (
	"booking-api/internal/domain/restaurant"
	"context"
)

type RestaurantRepository interface {
	Create(ctx context.Context, r *restaurant.Restaurant) error
	GetAll(ctx context.Context) ([]restaurant.Restaurant, error)
}
