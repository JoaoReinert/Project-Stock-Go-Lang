package domain

import (
	"Desafio_Go_Lang/entities"
	"context"
)

type EquipmentUseCase interface {
	RegisterEquipment(
		ctx context.Context,
		equipment entities.Equipment,
	) error

	UpdateEquipment(
		ctx context.Context,
		equipment entities.Equipment,
	) error

	DeleteEquipment(
		ctx context.Context,
		equipment entities.Equipment,
	) error

	DetailsEquipment(
		ctx context.Context,
		equipment entities.Equipment,
	) (*entities.Equipment, error)
}
