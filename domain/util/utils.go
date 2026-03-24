package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/scrypt"
	"io"
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
