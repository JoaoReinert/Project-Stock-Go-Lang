package datastore

import (
	"Desafio_Go_Lang/entities"
	"context"
)

type StockItemRepository interface {
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

	VerifyStock(
		ctx context.Context,
		stockItem entities.StockItem,
	) (int, error)
}
