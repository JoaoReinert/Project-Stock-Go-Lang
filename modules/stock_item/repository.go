package stock_item

import (
	"Desafio_Go_Lang/datastore"
	"Desafio_Go_Lang/entities"
	"context"
	"database/sql"
	"log"
)

type stockItemRepository struct {
	conn *sql.DB
}

func NewStockItemRepository(settings datastore.SettingsRepository) datastore.StockItemRepository {
	return stockItemRepository{
		conn: settings.Connection(),
	}
}

func (s stockItemRepository) AddStockItem(
	ctx context.Context,
	stockItem entities.StockItem,
) error {

	tx, err := s.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	//language=sql
	queryMovement := `
	INSERT INTO movement (date, id_user, type)
	VALUES (CURRENT_TIMESTAMP, ?, 0)
	`

	resultMovement, err := tx.ExecContext(
		ctx,
		queryMovement,
		7,
	)
	if err != nil {
		log.Printf("Error in [ExecContext], %v", err)
		_ = tx.Rollback()
		return err
	}

	idMovement, err := resultMovement.LastInsertId()
	if err != nil {
		log.Printf("Error in [LastInsertId], %v", err)
		_ = tx.Rollback()
		return err
	}
	//language=sql
	query := `
	INSERT INTO stock_item (id_equipment, id_unit_stock)
	VALUES (?, ?)
	`

	//language=sql
	queryHistoric := `
	INSERT INTO historic_movement (id_movement, id_stock_item)
	VALUES (?, ?)
	`

	quantity := int(stockItem.Quantity)

	for i := 0; i < quantity; i++ {
		result, err := tx.ExecContext(
			ctx,
			query,
			stockItem.IdEquipment,
			stockItem.IdUnitStock,
		)
		if err != nil {
			log.Printf("Error in [ExecContext], %v", err)
			_ = tx.Rollback()
			return err
		}

		id, err := result.LastInsertId()
		if err != nil {
			log.Printf("Error in [LastInsertId], %v", err)
			_ = tx.Rollback()
			return err
		}

		_, err = tx.ExecContext(
			ctx,
			queryHistoric,
			idMovement,
			id,
		)
		if err != nil {
			log.Printf("Error in [ExecContext], %v", err)
			_ = tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s stockItemRepository) RemoveStockItem(
	ctx context.Context,
	stockItem entities.StockItem,
) error {

	tx, err := s.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	//language=sql
	queryMovement := `
	INSERT INTO movement (date, id_user, type)
	VALUES (CURRENT_TIMESTAMP, ?, 1)
	`

	resultMovement, err := tx.ExecContext(
		ctx,
		queryMovement,
		7,
	)
	if err != nil {
		log.Printf("Error in [ExecContext], %v", err)
		_ = tx.Rollback()
		return err
	}
	idMovement, err := resultMovement.LastInsertId()
	if err != nil {
		log.Printf("Error in [LastInsertId], %v", err)
		_ = tx.Rollback()
		return err
	}
	//language=sql
	query := `
	INSERT INTO stock_item (id_equipment, id_unit_stock)
	values (?, ?)
	`
	//language=sql
	queryHistoric := `
	INSERT INTO historic_movement (id_movement, id_stock_item)
	VALUES (?, ?)
	`

	quantity := int(stockItem.Quantity)

	for i := 0; i < quantity; i++ {
		result, err := tx.ExecContext(
			ctx,
			query,
			stockItem.IdEquipment,
			stockItem.IdUnitStock,
		)
		if err != nil {
			log.Printf("Error in [ExecContext], %v", err)
			_ = tx.Rollback()
			return err
		}

		id, err := result.LastInsertId()
		if err != nil {
			log.Printf("Error in [LastInsertId], %v", err)
			_ = tx.Rollback()
			return err
		}

		_, err = tx.ExecContext(
			ctx,
			queryHistoric,
			idMovement,
			id,
		)
		if err != nil {
			log.Printf("Error in [ExecContext], %v", err)
			_ = tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (s stockItemRepository) VerifyStock(
	ctx context.Context,
	stockItem entities.StockItem,
) (int, error) {
	//language=sql
	query := `
	SELECT (
	    (SELECT COUNT(*) FROM stock_item 
	    WHERE id_equipment = ?
	    AND status_code = 0) 
	    -
	    (SELECT COUNT(*) FROM stock_item 
	    WHERE id_equipment = ?
	    AND status_code = 2)
	)
	`

	var number int
	err := s.conn.QueryRowContext(
		ctx,
		query,
		stockItem.IdEquipment,
		stockItem.IdEquipment,
	).Scan(&number)

	if err != nil {
		log.Printf("Error in [QueryRowContext]: %v", err)
		return 0, err
	}

	return number, nil
}
