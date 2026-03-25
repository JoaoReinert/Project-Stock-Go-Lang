package domain

import (
	"Desafio_Go_Lang/entities"
	"context"
)

type AuthenticationUseCase interface {
	RegisterUser(ctx context.Context, user entities.User) error

	CheckUserCredentials(ctx context.Context, user entities.UserLogin) (*entities.User, error)

	GenerateTokenUser(user entities.User) (*entities.UserToken, error)

	CheckDefaultSecurityToken(
		token entities.UserToken,
	) error
}
