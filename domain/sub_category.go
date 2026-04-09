package domain

import (
	"Desafio_Go_Lang/entities"
	"context"
)

type SubCategoryUseCase interface {
	RegisterSubCategory(
		ctx context.Context,
		category entities.SubCategory,
	) error

	UpdateSubCategory(
		ctx context.Context,
		category entities.SubCategory,
	) error

	DeleteSubCategory(
		ctx context.Context,
		category entities.SubCategory,
	) error

	DetailsSubCategory(
		ctx context.Context,
		subCategory entities.SubCategory,
	) (*entities.SubCategory, error)
}
