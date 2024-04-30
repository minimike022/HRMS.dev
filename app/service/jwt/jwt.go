package jwt

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
	//jwtware "github.com/gofiber/contrib/jwt"
)
var jwtSecret = "143"

func GenerateToken(user_name string, id int) (string, error) {
	claims := jwt.MapClaims{
		"id": id,
		"user_name": user_name,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString([]byte(jwtSecret))
}
