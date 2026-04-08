package postgres

import (
	"booking-api/internal/domain/restaurant"
	"booking-api/internal/repository"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RestaurantRepository struct {
	pool *pgxpool.Pool
}

func NewRestaurantRepository(pool *pgxpool.Pool) repository.RestaurantRepository {
	return &RestaurantRepository{pool: pool}
}
func (r *RestaurantRepository) Create(ctx context.Context, rest *restaurant.Restaurant) error {
	query := `INSERT INTO restaurants (name, address, phone) VALUES ($1, $2, $3) RETURNING id`
	err := r.pool.QueryRow(ctx, query, rest.Name, rest.Address, rest.Phone).Scan(&rest.ID)
	return err
}
func (r *RestaurantRepository) GetAll(ctx context.Context) ([]restaurant.Restaurant, error) {
	query := `SELECT id, name, address, phone FROM restaurants ORDER BY id`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var restaurants []restaurant.Restaurant
	for rows.Next() {
		var rest restaurant.Restaurant
		if err := rows.Scan(&rest.ID, &rest.Name, &rest.Address, &rest.Phone); err != nil {
			return nil, err
		}
		restaurants = append(restaurants, rest)
	}
	return restaurants, rows.Err()

}
func (r *RestaurantRepository) GetByID(ctx context.Context, id int64) (*restaurant.Restaurant, error) {
	var rest restaurant.Restaurant

	query := `SELECT id, name, address, phone FROM restaurants WHERE id = $1`

	row := r.pool.QueryRow(ctx, query, id)
	err := row.Scan(&rest.ID, &rest.Name, &rest.Address, &rest.Phone)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("restaurant not found")
		}
		return nil, err
	}
	return &rest, nil
}
func (r *RestaurantRepository) Update(ctx context.Context, rest *restaurant.Restaurant) error{
	query := `UPDATE restaurants SET name = $1, address = $2, phone = $3 WHERE id = $4`
	_, err := r.pool.Exec(ctx, query, rest.Name, rest.Address, rest.Phone, rest.ID)
	return err
}
func (r *RestaurantRepository) Delete(ctx context.Context, id int64) error{
	query := `DELETE FROM restaurants WHERE id = $1`
	result, err := r.pool.Exec(ctx, query, id)
	if err != nil{
		return err
	}
	if result.RowsAffected() == 0{
		return errors.New("restaurant not found")
	}
	return nil
}