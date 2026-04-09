package unitStock

import (
	"Desafio_Go_Lang/datastore"
	"Desafio_Go_Lang/domain"
	"Desafio_Go_Lang/entities"
	"context"
	"fmt"
	"log"
)

type unitStockUseCase struct {
	repository datastore.UnitStockRepository
	cfg        entities.Config
}

func NewUnitStockUseCase(repository datastore.UnitStockRepository, cfg entities.Config) domain.UnitStockUseCase {
	return unitStockUseCase{
		repository: repository,
		cfg:        cfg,
	}
}

func (u unitStockUseCase) RegisterUnitStock(
	ctx context.Context,
	unitStock entities.UnitStock,
) error {
	if unitStock.Name == "" {
		return fmt.Errorf("name is required")
	}

	err := u.repository.RegisterUnitStock(ctx, unitStock)
	if err != nil {
		log.Printf("Error in [RegisterUnitStock]")
		return err
	}

	return nil
}

func (u unitStockUseCase) UpdateUnitStock(
	ctx context.Context,
	unitStock entities.UnitStock,
) error {
	if unitStock.Name == "" {
		return fmt.Errorf("name is required")
	}

	err := u.repository.UpdateUnitStock(ctx, unitStock)
	if err != nil {
		log.Printf("Error in [UpdateUnitStock]")
		return err
	}

	return nil
}

func (u unitStockUseCase) DeleteUnitStock(
	ctx context.Context,
	unitStock entities.UnitStock,
) error {
	err := u.repository.DeleteUnitStock(ctx, unitStock)
	if err != nil {
		log.Printf("Error in [DeleteUnitStock]")
		return err
	}

	return nil
}

func (u unitStockUseCase) DetailsUnitStock(
	ctx context.Context,
	unitStock entities.UnitStock,
) (*entities.UnitStock, error) {
	details, err := u.repository.DetailsUnitStock(ctx, unitStock)
	if err != nil {
		log.Printf("Error in [DetailsUnitStock]")
		return nil, err
	}

	return details, nil
}
