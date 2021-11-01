// Package hasher provides functions to work with password hashing. SHA256 algorithm is being used.
package hasher

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

// HashPassword generates and returns a hash value for the given password
func HashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("Password is empty")
	}
	checksum := sha256.Sum256([]byte(password))
	hashedPass := string(base64.StdEncoding.EncodeToString(checksum[:]))
	return hashedPass, nil
}

// CheckPasswordHash checks if the given password matches the given hash value
func CheckPasswordHash(password, hash string) bool {
	hashedPass, err := HashPassword(password)
	if err != nil {
		return false
	}
	return hashedPass == hash
}
