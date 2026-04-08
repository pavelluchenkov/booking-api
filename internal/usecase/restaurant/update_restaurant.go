package restaurant

import (
	"booking-api/internal/domain/restaurant"
	"booking-api/internal/repository"
	"context"
	"errors"
	"strings"
)

type UpdateRestaurantUseCase struct {
	repo repository.RestaurantRepository
}

func NewUpdateRestaurant(repo repository.RestaurantRepository) *UpdateRestaurantUseCase {
	return &UpdateRestaurantUseCase{repo: repo}
}

func (uc *UpdateRestaurantUseCase) Execute(ctx context.Context, id int64, name, address, phone string) (*restaurant.Restaurant, error) {
	name = strings.TrimSpace(name)
	address = strings.TrimSpace(address)

	if name == "" {
		return nil, errors.New("name is required")
	}
	if address == ""{
		return nil, errors.New("address is required")
	}
	if phone == ""{
		return nil, errors.New("phone is required")
	}
	rest, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		if err.Error() == "restaurant not found" {
			return nil, errors.New("restaurant not found")
		}
		return nil, err
	}
	rest.Name = name
	rest.Address = address
	rest.Phone = strings.TrimSpace(phone)

	if err = uc.repo.Update(ctx, rest); err != nil{
		return nil, err
	}
	return rest, nil

}
