package equipment

import (
	"Desafio_Go_Lang/datastore"
	"Desafio_Go_Lang/entities"
	"context"
	"database/sql"
	"errors"
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
INSERT INTO equipment (name, id_sub_category, status_code)
values (?, ?, ?)
`

	_, err := e.conn.ExecContext(
		ctx, query,
		equipment.Name,
		equipment.IdSubCategory,
		0,
	)

	if err != nil {
		log.Printf("Error in [ExecContext]: %v", err)
		return err
	}

	return nil
}

func (e equipmentRepository) UpdateEquipment(
	ctx context.Context,
	equipment entities.Equipment,
) error {
	//language=sql
	query := `
	UPDATE equipment
	SET name = ?
		id_sub_category = ?
	WHERE id = ?
    `

	_, err := e.conn.ExecContext(ctx,
		query,
		equipment.Name,
		equipment.IdSubCategory,
		equipment.ID,
	)
	if err != nil {
		log.Printf("Error in [ExecContext]: %v", err)
		return err
	}

	return nil
}

func (e equipmentRepository) DeleteEquipment(
	ctx context.Context,
	equipment entities.Equipment,
) error {
	//language=sql
	query := `
	UPDATE equipment
	SET status_code = 2		
	WHERE id = ?
`
	_, err := e.conn.ExecContext(ctx, query, equipment.ID)
	if err != nil {
		log.Printf("Error in [ExecContext]: %v", err)
		return err
	}

	return nil
}

func (e equipmentRepository) DetailsEquipment(
	ctx context.Context,
	equipment entities.Equipment,
) (*entities.Equipment, error) {
	//language=sql
	query := `
	SELECT e.id,
		   e.name,	
		   e.id_sub_category,
		   e.status_code	
	FROM equipment e
	WHERE id = ?
`
	var details entities.Equipment

	err := e.conn.QueryRowContext(ctx, query, equipment.ID).Scan(
		&details.ID,
		&details.Name,
		&details.IdSubCategory,
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
