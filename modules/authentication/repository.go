package authentication

import (
	"Desafio_Go_Lang/datastore"
	"Desafio_Go_Lang/domain/util"
	"Desafio_Go_Lang/entities"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type authenticationRepository struct {
	conn *sql.DB
}

func NewAuthenticationRepository(settings datastore.SettingsRepository) datastore.AuthenticationRepository {
	return authenticationRepository{
		conn: settings.Connection(),
	}
}

func (e authenticationRepository) RegisterUser(
	ctx context.Context,
	user entities.User,
) error {
	// language=sql
	query := `
	INSERT INTO user (name, email, password, salt)
	VALUES (?, ?, ?, ?)
	`

	tx, err := e.conn.Begin()
	if err != nil {
		log.Printf("Error in [Beggin]")
		return err
	}

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error in [PrepareContext]")
		return err
	}

	_, err = stmt.ExecContext(ctx, user.Name, user.Email, user.Password, user.Salt)
	if err != nil {
		log.Printf("Error in [ExecContext]: %v", err)
		return tx.Rollback()
	}

	return tx.Commit()
}

func (e authenticationRepository) CheckUserCredentials(
	ctx context.Context,
	user entities.UserLogin,
) (*entities.User, error) {

	// language=sql
	query := `
	SELECT u.name,
	       u.email,
	       u.password,
	       u.salt
	FROM user u
	WHERE u.email = ?
	`

	var details entities.User
	var pass string
	var salt string

	err := e.conn.QueryRowContext(ctx, query, user.Email).Scan(
		&details.Name,
		&details.Email,
		&pass,
		&salt,
	)
	if err != nil {
		log.Printf("Error in [QueryRowContext]")
		return nil, err
	}

	hashedInput, err := util.SaltAndHash(user.Password, salt)
	if err != nil {
		log.Printf("Error in [SaltAndHash]: %v", err)
		return nil, err
	}

	if hashedInput == pass {
		return &details, nil
	}

	return nil, errors.New(fmt.Sprintf("user not found [%s]", user.Email))
}
