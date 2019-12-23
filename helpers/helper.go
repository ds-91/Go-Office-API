package helpers

import (
	"golang.org/x/crypto/bcrypt"
)


func IsEmpty(data string) bool {
	if len(data) == 0 {
		return true
	}

	return false
}

func HashPassword(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		panic("Error hashing password!")
	}
	return string(hash)
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}