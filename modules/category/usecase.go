package category

import (
	"Desafio_Go_Lang/datastore"
	"Desafio_Go_Lang/domain"
	"Desafio_Go_Lang/entities"
	"context"
	"fmt"
	"log"
)

type categoryUseCase struct {
	repository datastore.CategoryRepository
	cfg        entities.Config
}

func NewCategoryUseCase(repository datastore.CategoryRepository, cfg entities.Config) domain.CategoryUseCase {
	return categoryUseCase{
		repository: repository,
		cfg:        cfg,
	}
}

func (c categoryUseCase) RegisterCategory(
	ctx context.Context,
	category entities.Category,
) error {
	if category.Name == "" {
		return fmt.Errorf("name is required")
	}

	err := c.repository.RegisterCategory(ctx, category)
	if err != nil {
		log.Printf("Error in [RegisterCategory]")
		return err
	}

	return nil
}

func (c categoryUseCase) UpdateCategory(
	ctx context.Context,
	category entities.Category,
) error {
	if category.Name == "" {
		return fmt.Errorf("name is required")
	}

	err := c.repository.UpdateCategory(ctx, category)
	if err != nil {
		log.Printf("Error in [RegisterCategory]")
		return err
	}

	return nil
}

func (c categoryUseCase) DeleteCategory(
	ctx context.Context,
	category entities.Category,
) error {
	err := c.repository.DeleteCategory(ctx, category)
	if err != nil {
		log.Printf("Error in [RegisterCategory]")
		return err
	}

	return nil
}

func (c categoryUseCase) DetailsCategory(
	ctx context.Context,
	category entities.Category,
) (*entities.Category, error) {

	categoryDetails, err := c.repository.DetailsCategory(ctx, category)
	if err != nil {
		log.Printf("Error in [DetailsCategory]")
		return nil, err
	}

	return categoryDetails, nil
}
