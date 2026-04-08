package table

import (
	"booking-api/internal/domain/table"
	"booking-api/internal/repository"
	"context"
	"errors"
)

type GetTableByRestaurantID struct{
	tableRepo repository.TableRepository
	restaurantRepo repository.RestaurantRepository
}

func NewGetTableByRestaurantID(tableRepo repository.TableRepository, restaurantRepo repository.RestaurantRepository) *GetTableByRestaurantID{
	return &GetTableByRestaurantID{tableRepo: tableRepo, restaurantRepo: restaurantRepo}
}

func (uc *GetTableByRestaurantID) Execute(ctx context.Context, restaurantID int64) ([]table.Table, error){
	if restaurantID <= 0{
		return nil, errors.New("invalid restaurant_id: must be positive")
	}
	return uc.tableRepo.GetByRestaurantID(ctx, restaurantID)
}