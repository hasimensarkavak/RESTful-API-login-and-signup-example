package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPass), nil
}

func CheckPassword(password string, hasedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(password))
}
