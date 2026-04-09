package domain

import (
	"Desafio_Go_Lang/entities"
	"context"
)

type UnitStockUseCase interface {
	RegisterUnitStock(
		ctx context.Context,
		unitStock entities.UnitStock,
	) error

	UpdateUnitStock(
		ctx context.Context,
		unitStock entities.UnitStock,
	) error

	DeleteUnitStock(
		ctx context.Context,
		unitStock entities.UnitStock,
	) error

	DetailsUnitStock(
		ctx context.Context,
		unitStock entities.UnitStock,
	) (*entities.UnitStock, error)
}
