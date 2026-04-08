package restaurant

import (
	"booking-api/internal/repository"
	"context"
)

type DeleteRestaurantUseCase struct {
	repo repository.RestaurantRepository
}

func NewDeleteRestaurantUseCase(repo repository.RestaurantRepository) *DeleteRestaurantUseCase {
	return &DeleteRestaurantUseCase{repo: repo}
}
func (uc *DeleteRestaurantUseCase) Execute(ctx context.Context, id int64) error {
	_, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	return uc.repo.Delete(ctx, id)
}
