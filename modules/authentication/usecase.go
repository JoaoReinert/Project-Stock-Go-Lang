package authentication

import (
	"Desafio_Go_Lang/datastore"
	"Desafio_Go_Lang/domain"
	"Desafio_Go_Lang/domain/util"
	"Desafio_Go_Lang/entities"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/o1egl/paseto"
)

type authenticationUseCase struct {
	repository        datastore.AuthenticationRepository
	pasetoSecurityKey string
	passSaltSecret    string
}

func NewAuthenticationUseCase(repository datastore.AuthenticationRepository, pasetoSecurityKey string, passSaltSecret string) domain.AuthenticationUseCase {
	return authenticationUseCase{
		repository:        repository,
		pasetoSecurityKey: pasetoSecurityKey,
		passSaltSecret:    passSaltSecret,
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

	salt, err := util.GenerateSalt(e.passSaltSecret)
	if err != nil {
		log.Printf("Error in [util.GenerateSalt]: %v", err)
		return err
	}

	safePass, err := util.SaltAndHash(user.Password, salt)
	if err != nil {
		log.Printf("Error in [util.SaltAndHash]: %v", err)
		return err
	}

	newUser := entities.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: safePass,
		Salt:     salt,
	}

	return e.repository.RegisterUser(ctx, newUser)
}

func (e authenticationUseCase) CheckUserCredentials(ctx context.Context, user entities.UserLogin) (*entities.User, error) {

	if user.Email == "" {
		return nil, fmt.Errorf("email is required")
	}

	if user.Password == "" {
		return nil, fmt.Errorf("email is required")
	}

	if len(user.Password) < 6 {
		return nil, fmt.Errorf("the password must be longer than 6 characters")
	}

	userChecked, err := e.repository.CheckUserCredentials(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("error in [CheckUserCredentials] %v", err)
	}

	return userChecked, nil
}

func (e authenticationUseCase) GenerateTokenUser(user entities.User) (*entities.UserToken, error) {
	token, err := util.GenerateToken(user, e.pasetoSecurityKey)
	if err != nil {
		return nil, fmt.Errorf("Error in [GenerateToken]: %v", err)
	}

	return token, nil
}

func (e authenticationUseCase) CheckDefaultSecurityToken(
	token entities.UserToken,
) error {
	symmetricKey := []byte(e.pasetoSecurityKey)

	now := time.Now()

	var payload paseto.JSONToken
	var footer string

	_, err := paseto.Parse(token.Token, &payload, &footer, symmetricKey, nil)
	if err != nil {
		_ = fmt.Errorf("erro token: %s", err.Error())
	}

	if now.After(payload.Expiration) {
		_ = fmt.Errorf("token expirado")
	}

	return nil
}
