package postgres

import (
	"booking-api/internal/domain/table"
	"booking-api/internal/repository"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TableRepository struct {
	pool *pgxpool.Pool
}

func NewTableRepository(pool *pgxpool.Pool) repository.TableRepository {
	return &TableRepository{pool: pool}
}
func (r *TableRepository) Create(ctx context.Context, t *table.Table) error {
	query := `INSERT INTO tables (restaurants_id, number, capacity) VALUES ($1, $2, $3) RETURNING id`
	err := r.pool.QueryRow(ctx, query, t.RestaurantID, t.Number, t.Capacity).Scan(&t.ID)
	return err
}
func (r *TableRepository) GetByRestaurantID(ctx context.Context, restaurantID int64) ([]table.Table, error) {
	query := `SELECT id, restaurants_id, number, capacity, created_at FROM tables WHERE restaurants_id = $1 ORDER BY number`
	rows, err := r.pool.Query(ctx, query, restaurantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tables []table.Table
	for rows.Next() {
		var t table.Table
		if err := rows.Scan(&t.ID, &t.RestaurantID, &t.Number, &t.Capacity, &t.CreatedAt); err != nil {
			return nil, err
		}

		tables = append(tables, t)
	}
	return tables, rows.Err()
}
func (r *TableRepository) GetByID(ctx context.Context, id int64) (*table.Table, error) {
	var t table.Table
	query := `SELECT id, restaurants_id, number, capacity, created_at FROM tables WHERE id = $1`
	row := r.pool.QueryRow(ctx, query, id)
	err := row.Scan(&t.ID, &t.RestaurantID, &t.Number, &t.Capacity, &t.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("table not found")
		}
		return nil, err
	}
	return &t, nil
}
