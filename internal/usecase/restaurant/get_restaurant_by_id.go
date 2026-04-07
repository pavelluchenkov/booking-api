package restaurant

import (
	"booking-api/internal/domain/restaurant"
	"booking-api/internal/repository"
	"context"
	"errors"
)

type GetRestaurantByIDUseCase struct {
	repo repository.RestaurantRepository
}

func NewGetRestaurantByIDUseCase(repo repository.RestaurantRepository) *GetRestaurantByIDUseCase {
	return &GetRestaurantByIDUseCase{repo: repo}
}

func (uc *GetRestaurantByIDUseCase) Execute(ctx context.Context, id int64) (*restaurant.Restaurant, error) {
	if id <= 0 {
		return nil, errors.New("invalid id: must be positive")
	}
	return uc.repo.GetByID(ctx, id)
}
