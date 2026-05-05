package statistics

import (
	"Desafio_Go_Lang/datastore"
	"Desafio_Go_Lang/domain"
	"Desafio_Go_Lang/entities"
	"context"
	"log"
	"time"
)

type statisticsUseCase struct {
	repository datastore.StatisticsRepository
	cfg        entities.Config
}

func NewStatisticsUseCase(repository datastore.StatisticsRepository, cfg entities.Config) domain.StatisticsUseCase {
	return statisticsUseCase{
		repository: repository,
		cfg:        cfg,
	}
}

func (s statisticsUseCase) GetTotalNumberPerDate(
	ctx context.Context,
	startDate time.Time,
	endDate time.Time,
	isEntries bool,
) (int, error) {

	total, err := s.repository.GetTotalNumberPerDate(ctx, startDate, endDate, isEntries)
	if err != nil {
		log.Printf("Error in [GetTotalNumberEntriesPerDate]")
		return 0, err
	}

	return total, nil
}
