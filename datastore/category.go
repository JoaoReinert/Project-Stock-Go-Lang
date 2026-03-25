package datastore

import (
	"Desafio_Go_Lang/entities"
	"context"
)

type CategoryRepository interface {
	RegisterCategory(
		ctx context.Context,
		category entities.Category,
	) error

	UpdateCategory(
		ctx context.Context,
		category entities.Category,
	) error

	DeleteCategory(
		ctx context.Context,
		category entities.Category,
	) error

	DetailsCategory(
		ctx context.Context,
		category entities.Category,
	) (*entities.Category, error)
}
