package restaurant

import (
	"booking-api/internal/domain/restaurant"
	"booking-api/internal/repository"
	"context"
	"errors"
	"strings"
)

type CreateRestaurantUseCase struct {
	repo repository.RestaurantRepository
}

func NewCreateRestaurantUseCase(repo repository.RestaurantRepository) *CreateRestaurantUseCase {
	return &CreateRestaurantUseCase{repo: repo}
}

func (uc *CreateRestaurantUseCase) Execute(ctx context.Context, name, address, phone string) (*restaurant.Restaurant, error){
	name = strings.TrimSpace(name)
	address = strings.TrimSpace(address)

	if name == ""{
		return nil, errors.New("name is required")
	}
	if address == ""{
		return nil, errors.New("address if required")
	}

	rest := &restaurant.Restaurant{
		Name: name,
		Address: address,
		Phone: strings.TrimSpace(phone),
	}
	if err := uc.repo.Create(ctx, rest); err != nil{
		return nil, err
	}
	return rest, nil
}
