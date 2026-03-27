package equipment

import (
	"Desafio_Go_Lang/datastore"
	"Desafio_Go_Lang/entities"
	"context"
	"database/sql"
	"log"
)

type equipmentRepository struct {
	conn *sql.DB
}

func NewEquipmentRepository(settings datastore.SettingsRepository) datastore.EquipmentRepository {
	return equipmentRepository{
		conn: settings.Connection(),
	}
}

func (e equipmentRepository) RegisterEquipment(
	ctx context.Context,
	equipment entities.Equipment,
) error {
	//language=sql
	query := `
	INSERT INTO equipment (name, id_sub_category)  
	VALUES (?, ?)
`
	_, err := e.conn.ExecContext(ctx, query, equipment.Name, equipment.IdSubCategory)
	if err != nil {
		log.Printf("Error in [ExecContext]: %v", err)
		return err
	}

	return nil
}
