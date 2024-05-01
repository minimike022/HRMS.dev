package util

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)
var jwtSecret = "143"

func JwtAuth() func(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(jwtSecret)},
	})
} 

func HashedPassword(password string) string {
	hashed_password, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed_password)
}

func CompareHash(stored_password string, forms_password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(stored_password), []byte(forms_password))
	return err
}

