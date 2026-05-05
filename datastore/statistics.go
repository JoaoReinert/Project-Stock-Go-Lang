package datastore

import (
	"Desafio_Go_Lang/entities"
	"context"
	"time"
)

type StatisticsRepository interface {
	GetTotalNumberPerDate(
		ctx context.Context,
		startDate time.Time,
		endDate time.Time,
		isEntries bool,
	) (int, error)

	GetBalancePerUnitStock(
		ctx context.Context,
	) ([]entities.UnitStockBalance, error)
}
