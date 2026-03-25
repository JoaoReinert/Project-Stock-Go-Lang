package category

import (
	"Desafio_Go_Lang/datastore"
	"Desafio_Go_Lang/entities"
	"context"
	"database/sql"
	"log"
)

type categoryRepository struct {
	conn *sql.DB
}

func NewCategoryRepository(settings datastore.SettingsRepository) datastore.CategoryRepository {
	return categoryRepository{
		conn: settings.Connection(),
	}
}

func (c categoryRepository) RegisterCategory(
	ctx context.Context,
	category entities.Category,
) error {
	//language=sql
	query := `
	INSERT INTO category (name)
	VALUES (?)
	`

	_, err := c.conn.ExecContext(ctx, query, category.Name)
	if err != nil {
		log.Printf("Error in [ExecContext]: %v", err)
		return err
	}

	return nil
}

func (c categoryRepository) UpdateCategory(
	ctx context.Context,
	category entities.Category,
) error {
	//language=sql
	query := `
	UPDATE category
	SET name = ?
	WHERE id = ?
    `

	_, err := c.conn.ExecContext(ctx, query, category.Name, category.ID)
	if err != nil {
		log.Printf("Error in [ExecContext]: %v", err)
		return err
	}

	return nil
}

func (c categoryRepository) DeleteCategory(
	ctx context.Context,
	category entities.Category,
) error {
	//language=sql
	query := `
	DELETE FROM category
	WHERE id = ?
`
	_, err := c.conn.ExecContext(ctx, query, category.ID)
	if err != nil {
		log.Printf("Error in [ExecContext]: %v", err)
		return err
	}

	return nil
}

func (c categoryRepository) DetailsCategory(
	ctx context.Context,
	category entities.Category,
) (*entities.Category, error) {
	//language=sql
	query := `
	SELECT c.id,
		   c.name	
	FROM category c
	WHERE id = ?
`
	var details entities.Category

	err := c.conn.QueryRowContext(ctx, query, category.ID).Scan(
		&details.ID,
		&details.Name,
	)
	if err != nil {
		log.Printf("Error in [QueryRowContext]")
		return nil, err
	}

	return &details, nil
}
