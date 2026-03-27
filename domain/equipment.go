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
}
