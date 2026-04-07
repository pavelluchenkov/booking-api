package postgres

import (
	"booking-api/internal/domain/restaurant"
	"booking-api/internal/repository"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RestaurantRepository struct{
	pool *pgxpool.Pool
}
func NewRestaurantRepository(pool *pgxpool.Pool) repository.RestaurantRepository{
	return &RestaurantRepository{pool: pool}
}
func (r *RestaurantRepository) Create(ctx context.Context, rest *restaurant.Restaurant) error{
	query := `INSERT INTO restaurants (name, address, phone) VALUES ($1, $2, $3) RETURNING id`
	err := r.pool.QueryRow(ctx, query, rest.Name, rest.Address, rest.Phone).Scan(&rest.ID)
	return err
}