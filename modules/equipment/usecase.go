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

func (c equipmentUseCase) RegisterEquipment(
	ctx context.Context,
	equipment entities.Equipment,
) error {
	if equipment.Name == "" {
		return fmt.Errorf("name is required")
	}

	err := c.repository.RegisterEquipment(ctx, equipment)
	if err != nil {
		log.Printf("Error in [RegisterCategory]")
		return err
	}

	return nil
}
