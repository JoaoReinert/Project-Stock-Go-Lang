package domain

import (
	"Desafio_Go_Lang/entities"
	"context"
)

type UserUseCase interface {
	RegisterUser(ctx context.Context, user entities.User) error
}
