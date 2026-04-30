package util

import (
	"Desafio_Go_Lang/entities"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/scrypt"
)

func SaltAndHash(informationToHash string, salt string) (string, error) {
	generated, err := scrypt.Key(
		[]byte(informationToHash),
		[]byte(salt),
		1<<16,
		8,
		1,
		64,
	)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(generated), nil
}

func GenerateSalt(passSaltSecret string) (string, error) {
	secret := passSaltSecret
	buf := make([]byte, 32, 32+sha256.Size)
	_, err := io.ReadFull(rand.Reader, buf)

	if err != nil {
		return "", err
	}

	hash := sha256.New()
	hash.Write(buf)
	hash.Write([]byte(secret))

	return hex.EncodeToString(hash.Sum(buf)), nil
}

func GenerateToken(user entities.User, pasetoSecurityKey string) (*entities.UserToken, error) {
	v2 := paseto.NewV2()
	now := time.Now()
	expiration := now.Add(24 * time.Hour)
	symmetricKey := []byte(pasetoSecurityKey)

	tokenUuid, err := uuid.NewRandom()
	if err != nil {
		log.Printf("Error in [NewRandom]: %v", err)
		return nil, fmt.Errorf("error in generate uuid of token: %s", err.Error())
	}

	ts := entities.EmployeeSubject{
		ID: &user.ID,
	}

	js, err := json.Marshal(ts)
	if err != nil {
		log.Printf("Error in [Marshal]: %v", err)
		return nil, err
	}

	idString := strconv.FormatInt(user.ID, 10)

	jsonToken := paseto.JSONToken{
		Audience:   idString,
		Issuer:     "/api",
		Jti:        tokenUuid.String(),
		Subject:    string(js),
		Expiration: expiration,
		IssuedAt:   now,
		NotBefore:  now,
	}

	encripted, err := v2.Encrypt(symmetricKey, jsonToken, "")
	if err != nil {
		log.Printf("Error in [Encrypt]: %v", err)
		return nil, fmt.Errorf("error in encrypting token: %s", err.Error())
	}

	token := &entities.UserToken{
		Token: encripted,
	}

	return token, nil
}

func GetUser(r *http.Request) (*entities.User, error) {
	contextUser := r.Context().Value("user")
	if contextUser == nil {
		return nil, errors.New("user not found in the request")
	}

	return contextUser.(*entities.User), nil
}
