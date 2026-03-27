package datastore

import (
	"Desafio_Go_Lang/entities"
	"context"
)

type EquipmentRepository interface {
	RegisterEquipment(
		ctx context.Context,
		equipment entities.Equipment,
	) error
}
