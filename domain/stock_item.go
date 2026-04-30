package domain

import (
	"Desafio_Go_Lang/entities"
	"context"
)

type StockItemUseCase interface {
	AddStockItem(
		ctx context.Context,
		stockItem entities.StockItem,
		userId int64,
	) error

	RemoveStockItem(
		ctx context.Context,
		stockItem entities.StockItem,
		userId int64,
	) error
}
