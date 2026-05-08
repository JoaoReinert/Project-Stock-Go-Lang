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

func (s statisticsUseCase) GetBalancePerUnitStock(
	ctx context.Context,
) ([]entities.UnitStockBalance, error) {

	list, err := s.repository.GetBalancePerUnitStock(ctx)
	if err != nil {
		log.Printf("Error in [GetBalancePerUnitStock]")
		return nil, err
	}

	return list, nil
}

func (s statisticsUseCase) GetBalancePerCategoryAndSubCategory(
	ctx context.Context,
) ([]entities.CategoryBalance, error) {
	list, err := s.repository.GetBalancePerCategoryAndSubCategory(ctx)
	if err != nil {
		log.Printf("Error in [GetBalancePerCategoryAndSubCategory]")
		return nil, err
	}

	return list, nil
}
