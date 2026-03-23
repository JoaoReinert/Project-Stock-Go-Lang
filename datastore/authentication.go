package datastore

import (
	"Desafio_Go_Lang/entities"
	"context"
)

type AuthenticationRepository interface {
	RegisterUser(
		ctx context.Context,
		user entities.User,
	) error
}
