package statistics

import (
	"Desafio_Go_Lang/datastore"
	"Desafio_Go_Lang/entities"
	"context"
	"database/sql"
	"time"
)

type statisticsRepository struct {
	conn *sql.DB
}

func NewStatisticsRepository(settings datastore.SettingsRepository) datastore.StatisticsRepository {
	return statisticsRepository{
		conn: settings.Connection(),
	}
}

func (s statisticsRepository) GetTotalNumberPerDate(
	ctx context.Context,
	startDate time.Time,
	endDate time.Time,
	isExist bool,
) (int, error) {
	//language=sql
	query := `
	SELECT COUNT(hm.id)
	FROM historic_movement hm
	INNER JOIN movement m ON m.id = hm.id_movement
	WHERE m.type = ?
		AND m.date BETWEEN DATE(?) AND DATE(?);
	`

	var total int

	err := s.conn.QueryRowContext(
		ctx,
		query,
		isExist,
		startDate,
		endDate,
	).Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
}

func (s statisticsRepository) GetBalancePerUnitStock(
	ctx context.Context,
) ([]entities.UnitStockBalance, error) {

	//language=sql
	query := `
	SELECT us.name,
	       SUM(IF(m.type = 0, 1, -1)) AS balance
	FROM unit_stock us
	INNER JOIN stock_item si ON si.id_unit_stock = us.id
	INNER JOIN historic_movement hm ON hm.id_stock_item = si.id
	INNER JOIN movement m ON m.id = hm.id_movement
	GROUP BY us.id, us.name
	`

	rows, err := s.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []entities.UnitStockBalance

	for rows.Next() {
		var item entities.UnitStockBalance

		err := rows.Scan(&item.Name, &item.Balance)
		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	return result, nil
}
