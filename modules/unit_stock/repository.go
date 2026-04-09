package unitStock

import (
	"Desafio_Go_Lang/datastore"
	"Desafio_Go_Lang/entities"
	"context"
	"database/sql"
	"errors"
	"log"
)

type unitStockRepository struct {
	conn *sql.DB
}

func NewUnitStockRepository(settings datastore.SettingsRepository) datastore.UnitStockRepository {
	return unitStockRepository{
		conn: settings.Connection(),
	}
}

func (u unitStockRepository) RegisterUnitStock(
	ctx context.Context,
	unitStock entities.UnitStock,
) error {
	//language=sql
	query := `
	INSERT INTO unit_stock (name, status_code)
	values (?, ?)
	`

	_, err := u.conn.ExecContext(
		ctx,
		query,
		unitStock.Name,
		0,
	)

	if err != nil {
		log.Printf("Error in [ExecContext]: %v", err)
		return err
	}

	return nil
}

func (u unitStockRepository) UpdateUnitStock(
	ctx context.Context,
	unitStock entities.UnitStock,
) error {
	//language=sql
	query := `
	UPDATE unit_stock
	SET   name = ?
	WHERE id = ?
	`

	_, err := u.conn.ExecContext(ctx, query, unitStock.Name, unitStock.ID)
	if err != nil {
		log.Printf("Error in [ExecContext]: %v", err)
		return err
	}

	return nil
}

func (u unitStockRepository) DeleteUnitStock(
	ctx context.Context,
	unitStock entities.UnitStock,
) error {
	//language=sql
	query := `
	UPDATE unit_stock
	SET status_code = 2
	WHERE id = ?
	`

	_, err := u.conn.ExecContext(ctx, query, unitStock.ID)
	if err != nil {
		log.Printf("Error in [ExecContext]: %v", err)
		return err
	}

	return nil
}

func (u unitStockRepository) DetailsUnitStock(
	ctx context.Context,
	unitStock entities.UnitStock,
) (*entities.UnitStock, error) {
	//language=sql
	query := `
	SELECT u.id,
	       u.name,
	       u.status_code
	FROM unit_stock u   
	WHERE id = ?
	`

	var details entities.UnitStock

	err := u.conn.QueryRowContext(ctx, query, unitStock.ID).Scan(
		&details.ID,
		&details.Name,
		&details.StatusCode,
	)

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Printf("Error in [QueryRowContext]")
			return nil, err
		}

		return nil, nil
	}

	return &details, nil
}
