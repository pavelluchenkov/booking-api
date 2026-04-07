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
func (r *RestaurantRepository) GetAll(ctx context.Context) ([]restaurant.Restaurant, error){
	query := `SELECT id, name, address, phone FROM restaurants ORDER BY id`
	rows, err := r.pool.Query(ctx, query)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var restaurants []restaurant.Restaurant
	for rows.Next(){
		var rest restaurant.Restaurant
		if err := rows.Scan(&rest.ID, &rest.Name, &rest.Address, &rest.Phone); err != nil{
			return nil, err
		}
		restaurants = append(restaurants, rest)
	}
	return restaurants, rows.Err()
	

}