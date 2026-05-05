package statistics

import (
	"Desafio_Go_Lang/datastore"
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
	isEntries bool,
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
		isEntries,
		startDate,
		endDate,
	).Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
}
