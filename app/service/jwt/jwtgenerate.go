package jwtgenerate

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
)


func GenerateRefreshToken(user_name string, id int, user_role string) (string, error) {
	claims := jwt.MapClaims{
		"id": id,
		"user_name": user_name,
		"user_role": user_role,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString([]byte(os.Getenv("REFRESH_TOKEN_KEY")))
}

func GenerateAccessToken(user_name string, id int, user_role string) (string, error) {
	claims := jwt.MapClaims{
		"id": id,
		"user_name": user_name,
		"user_role": user_role,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString([]byte(os.Getenv("REFRESH_ACCESS_KEY")))
}




