package sub_category

import (
	"Desafio_Go_Lang/datastore"
	"Desafio_Go_Lang/entities"
	"context"
	"database/sql"
	"errors"
	"log"
)

type subCategoryRepository struct {
	conn *sql.DB
}

func NewSubCategoryRepository(settings datastore.SettingsRepository) datastore.SubCategoryRepository {
	return subCategoryRepository{
		conn: settings.Connection(),
	}
}

func (c subCategoryRepository) RegisterSubCategory(
	ctx context.Context,
	subCategory entities.SubCategory,
) error {
	//language=sql
	query := `
	INSERT INTO sub_category (name, id_category)
	VALUES (?, ?)
`
	_, err := c.conn.ExecContext(ctx, query, subCategory.Name, subCategory.IdCategory)
	if err != nil {
		log.Printf("Error in [ExecContext]: %v", err)
		return err
	}

	return nil
}

func (c subCategoryRepository) UpdateSubCategory(
	ctx context.Context,
	subCategory entities.SubCategory,
) error {
	//language=sql
	query := `
	UPDATE sub_category
	SET name = ?,
		id_category = ?
	WHERE id = ?
`
	_, err := c.conn.ExecContext(ctx, query, subCategory.Name, subCategory.IdCategory, subCategory.ID)
	if err != nil {
		log.Printf("Error in [ExecContext]: %v", err)
		return err
	}

	return nil
}

func (c subCategoryRepository) DeleteSubCategory(
	ctx context.Context,
	subCategory entities.SubCategory,
) error {
	//language=sql
	query := `
	UPDATE sub_category
	SET status_code = 2		
	WHERE id = ?
`

	_, err := c.conn.ExecContext(ctx, query, subCategory.ID)
	if err != nil {
		log.Printf("Error in [ExecContext] %v", err)
		return err
	}

	return nil
}

func (c subCategoryRepository) DetailsSubCategory(
	ctx context.Context,
	subCategory entities.SubCategory,
) (*entities.SubCategory, error) {
	//language=sql
	query := `
	SELECT s.id,
		   s.name,	
		   s.id_category	
	FROM sub_category s
	WHERE id = ?
`
	var details entities.SubCategory

	err := c.conn.QueryRowContext(ctx, query, subCategory.ID).Scan(
		&details.ID,
		&details.Name,
		&details.IdCategory,
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
