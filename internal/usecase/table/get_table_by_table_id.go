package table

import (
	"booking-api/internal/domain/table"
	"booking-api/internal/repository"
	"context"
	"errors"
)

type GetTableByTableID struct{
	tableRepo repository.TableRepository
}

func NewGetTableByTableID(tableRepo repository.TableRepository) *GetTableByTableID{
	return &GetTableByTableID{tableRepo: tableRepo}
}

func (uc *GetTableByTableID) Execute(ctx context.Context, id int64) (*table.Table, error){
	if id <= 0{
		return nil, errors.New("invalid id: must be positive")
	}
	return uc.tableRepo.GetByID(ctx, id)
}