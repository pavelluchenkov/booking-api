package restaurant

import (
	"booking-api/internal/domain/restaurant"
	"booking-api/internal/repository"
	"context"
)

type GetAllRestaurantsUseCase struct {
	repo repository.RestaurantRepository
}

func NewGetAllRestaurantsUseCase(repo repository.RestaurantRepository) *GetAllRestaurantsUseCase {
	return &GetAllRestaurantsUseCase{repo: repo}
}

func (uc *GetAllRestaurantsUseCase) Execute(ctx context.Context) ([]restaurant.Restaurant, error) {
	return uc.repo.GetAll(ctx)
}
