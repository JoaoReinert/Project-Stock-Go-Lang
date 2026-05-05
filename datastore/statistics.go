package datastore

import (
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
}
