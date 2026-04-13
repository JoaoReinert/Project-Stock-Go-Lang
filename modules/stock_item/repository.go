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
	query := `
	INSERT INTO stock_item (id_equipment, id_unit_stock, type_movement)
	values (?, ?, ?)
	`

	quantity := int(stockItem.Quantity)

	for i := 0; i < quantity; i++ {
		_, err := tx.ExecContext(
			ctx,
			query,
			stockItem.IdEquipment,
			stockItem.IdUnitStock,
			0,
		)

		if err != nil {
			log.Printf("Error in [ExecContext], %v", err)
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
	query := `
	INSERT INTO stock_item (id_equipment, id_unit_stock, type_movement)
	values (?, ?, ?)
	`

	quantity := int(stockItem.Quantity)

	for i := 0; i < quantity; i++ {
		_, err := tx.ExecContext(
			ctx,
			query,
			stockItem.IdEquipment,
			stockItem.IdUnitStock,
			1,
		)

		if err != nil {
			log.Printf("Error in [ExecContext], %v", err)
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
	    AND type_movement = 0) 
	    -
	    (SELECT COUNT(*) FROM stock_item 
	    WHERE id_equipment = ?
	    AND type_movement = 1)
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

func (s stockItemRepository) VerifyQuantityItemsStock(
	ctx context.Context,
	stockItem entities.StockItem,
) (int, error) {
	//language=sql
	query := `
	    SELECT COUNT(*) FROM stock_item 
	    WHERE id_equipment = ?
	    AND type_movement = 0
	`

	var number int
	err := s.conn.QueryRowContext(
		ctx,
		query,
		stockItem.IdEquipment,
	).Scan(&number)

	if err != nil {
		log.Printf("Error in [QueryRowContext]: %v", err)
		return 0, err
	}

	return number, nil
}
