package domain

import (
	"context"
	"time"
)

type StatisticsUseCase interface {
	GetTotalNumberPerDate(
		ctx context.Context,
		startDate time.Time,
		endDate time.Time,
		isEntries bool,
	) (int, error)
}
