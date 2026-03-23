package authentication

import (
	"Desafio_Go_Lang/datastore"
	"Desafio_Go_Lang/domain"
	"Desafio_Go_Lang/entities"
	"context"
	"fmt"
)

type authenticationUseCase struct {
	repository        datastore.AuthenticationRepository
	pasetoSecurityKey string
}

func NewAuthenticationUseCase(repository datastore.AuthenticationRepository, pasetoSecurityKey string) domain.AuthenticationUseCase {
	return authenticationUseCase{
		repository:        repository,
		pasetoSecurityKey: pasetoSecurityKey,
	}
}

func (e authenticationUseCase) RegisterUser(ctx context.Context, user entities.User) error {

	if user.Name == "" {
		return fmt.Errorf("name is required")
	}

	if user.Email == "" {
		return fmt.Errorf("email is required")
	}

	if user.Password == "" {
		return fmt.Errorf("password is required")
	}

	if len(user.Password) < 6 {
		return fmt.Errorf("the password must be longer than 6 characters")
	}

	return e.repository.RegisterUser(ctx, user)
}
