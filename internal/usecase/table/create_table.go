package table

import (
	"booking-api/internal/domain/table"
	"booking-api/internal/repository"
	"context"
	"errors"
)

type CreateTableUseCase struct {
	tableRepo      repository.TableRepository
	restaurantRepo repository.RestaurantRepository
}

func NewCreateTableUseCase(tableRepo repository.TableRepository, restaurantRepo repository.RestaurantRepository) *CreateTableUseCase {
	return &CreateTableUseCase{tableRepo: tableRepo, restaurantRepo: restaurantRepo}
}

func (uc *CreateTableUseCase) Execute(ctx context.Context, restaurantID int64, number, capacity int) (*table.Table, error) {
	if restaurantID <= 0 {
		return nil, errors.New("invalid restaurant id")
	}
	if number <= 0 {
		return nil, errors.New("table number must be positive")
	}
	if capacity <= 0 {
		return nil, errors.New("capacity must be positive")
	}
	_, err := uc.restaurantRepo.GetByID(ctx, restaurantID)
	if err != nil{
		return nil, err
	}
	t := &table.Table{
		RestaurantID: restaurantID,
		Number: number,
		Capacity: capacity,
	}
	if err := uc.tableRepo.Create(ctx, t); err != nil{
		return nil, err
	}
	return t, nil
}
