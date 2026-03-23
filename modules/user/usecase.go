package user

import (
	"Desafio_Go_Lang/datastore"
	"Desafio_Go_Lang/domain"
	"Desafio_Go_Lang/entities"
	"context"
	"fmt"
)

type userUseCase struct {
	repository datastore.UserRepository
}

func NewUserUseCase(repository datastore.UserRepository) domain.UserUseCase {
	return userUseCase{repository: repository}
}

func (e userUseCase) RegisterUser(ctx context.Context, user entities.User) error {

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
