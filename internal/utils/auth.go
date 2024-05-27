package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes),err
}

func ComparePassword(userPassword, enteredPasswrod string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(enteredPasswrod))
	return err
} 