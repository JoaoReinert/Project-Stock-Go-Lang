package domain

import (
	"Desafio_Go_Lang/entities"
	"context"
)

type AuthenticationUseCase interface {
	RegisterUser(ctx context.Context, user entities.User) error
}
