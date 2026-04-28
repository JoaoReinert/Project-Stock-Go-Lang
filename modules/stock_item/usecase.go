package stock_item

import (
	"Desafio_Go_Lang/datastore"
	"Desafio_Go_Lang/domain"
	"Desafio_Go_Lang/entities"
	"context"
	"fmt"
	"log"
)

type stockItemUseCase struct {
	repository datastore.StockItemRepository
	cfg        entities.Config
}

func NewStockItemUseCase(repository datastore.StockItemRepository, cfg entities.Config) domain.StockItemUseCase {
	return stockItemUseCase{
		repository: repository,
		cfg:        cfg,
	}
}

func (u stockItemUseCase) AddStockItem(
	ctx context.Context,
	stockItem entities.StockItem,
) error {
	if stockItem.IdEquipment == 0 {
		return fmt.Errorf("id equipment required")
	}

	if stockItem.IdUnitStock == 0 {
		return fmt.Errorf("id unit stock required")
	}

	if stockItem.Quantity == 0 {
		return fmt.Errorf("quantity required")
	}

	err := u.repository.AddStockItem(ctx, stockItem)
	if err != nil {
		log.Printf("Error in [RegisterStockItem]")
		return err
	}

	return nil
}

func (u stockItemUseCase) RemoveStockItem(
	ctx context.Context,
	stockItem entities.StockItem,
) error {
	if stockItem.IdEquipment == 0 {
		return fmt.Errorf("id equipment required")
	}

	if stockItem.IdUnitStock == 0 {
		return fmt.Errorf("id unit stock required")
	}

	if stockItem.Quantity == 0 {
		return fmt.Errorf("quantity required")
	}

	stock, errStock := u.repository.VerifyStock(ctx, stockItem)
	if errStock != nil {
		return errStock
	}

	if stock < int(stockItem.Quantity) {
		return fmt.Errorf("insufficient stock")
	}

	err := u.repository.RemoveStockItem(ctx, stockItem)
	if err != nil {
		log.Printf("Error in [RemoveStockItem]")
		return err
	}

	return nil
}
