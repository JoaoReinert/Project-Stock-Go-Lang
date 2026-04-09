package equipment

import (
	"Desafio_Go_Lang/datastore"
	"Desafio_Go_Lang/domain"
	"Desafio_Go_Lang/entities"
	"context"
	"fmt"
	"log"
)

type equipmentUseCase struct {
	repository datastore.EquipmentRepository
	cfg        entities.Config
}

func NewEquipmentUseCase(repository datastore.EquipmentRepository, cfg entities.Config) domain.EquipmentUseCase {
	return equipmentUseCase{
		repository: repository,
		cfg:        cfg,
	}
}

func (e equipmentUseCase) RegisterEquipment(
	ctx context.Context,
	equipment entities.Equipment,
) error {
	if equipment.Name == "" {
		return fmt.Errorf("name is required")
	}

	if equipment.IdSubCategory == 0 {
		return fmt.Errorf("sub category is required")
	}

	err := e.repository.RegisterEquipment(ctx, equipment)
	if err != nil {
		log.Printf("Error in [RegisterEquipment]")
		return err
	}

	return nil
}

func (e equipmentUseCase) UpdateEquipment(
	ctx context.Context,
	equipment entities.Equipment,
) error {
	if equipment.Name == "" {
		return fmt.Errorf("name is required")
	}

	if equipment.IdSubCategory == 0 {
		return fmt.Errorf("sub category is required")
	}

	err := e.repository.UpdateEquipment(ctx, equipment)
	if err != nil {
		log.Printf("Error in [UpdateEquipment]")
		return err
	}

	return nil
}

func (e equipmentUseCase) DeleteEquipment(
	ctx context.Context,
	equipment entities.Equipment,
) error {
	err := e.repository.DeleteEquipment(ctx, equipment)
	if err != nil {
		log.Printf("Error in [DeleteEquipment]")
		return err
	}

	return nil
}

func (e equipmentUseCase) DetailsEquipment(
	ctx context.Context,
	equipment entities.Equipment,
) (*entities.Equipment, error) {

	equipmentDetails, err := e.repository.DetailsEquipment(ctx, equipment)
	if err != nil {
		log.Printf("Error in [DetailsEquipment]")
		return nil, err
	}

	return equipmentDetails, nil
}
