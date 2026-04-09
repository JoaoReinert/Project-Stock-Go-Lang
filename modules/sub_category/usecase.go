package sub_category

import (
	"Desafio_Go_Lang/datastore"
	"Desafio_Go_Lang/domain"
	"Desafio_Go_Lang/entities"
	"context"
	"fmt"
	"log"
)

type subCategoryUseCase struct {
	repository datastore.SubCategoryRepository
	cfg        entities.Config
}

func NewSubCategoryUseCase(repository datastore.SubCategoryRepository, cfg entities.Config) domain.SubCategoryUseCase {
	return subCategoryUseCase{
		repository: repository,
		cfg:        cfg,
	}
}

func (c subCategoryUseCase) RegisterSubCategory(
	ctx context.Context,
	subCategory entities.SubCategory,
) error {
	if subCategory.Name == "" {
		return fmt.Errorf("name is required")
	}

	err := c.repository.RegisterSubCategory(ctx, subCategory)
	if err != nil {
		log.Printf("Error in [RegisterSubCategory]")
		return err
	}

	return nil
}

func (c subCategoryUseCase) UpdateSubCategory(
	ctx context.Context,
	subCategory entities.SubCategory,
) error {
	if subCategory.Name == "" {
		return fmt.Errorf("name is required")
	}

	err := c.repository.UpdateSubCategory(ctx, subCategory)
	if err != nil {
		log.Printf("Error in [UpdateSubCategory]")
		return err
	}

	return nil
}

func (c subCategoryUseCase) DeleteSubCategory(
	ctx context.Context,
	subCategory entities.SubCategory,
) error {
	err := c.repository.DeleteSubCategory(ctx, subCategory)
	if err != nil {
		log.Printf("Error in [DeleteSubCategory]")
		return err
	}

	return nil
}

func (c subCategoryUseCase) DetailsSubCategory(
	ctx context.Context,
	subCategory entities.SubCategory,
) (*entities.SubCategory, error) {

	subCategoryDetails, err := c.repository.DetailsSubCategory(ctx, subCategory)
	if err != nil {
		log.Printf("Error in [DetailsSubCategory]")
		return nil, err
	}

	return subCategoryDetails, nil
}
