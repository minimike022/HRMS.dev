package jwt

import (
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)
var jwtSecret = "143"

func JwtAuth() func(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func (ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized Access",
			})
		},
		SigningKey: jwtware.SigningKey{Key: []byte(jwtSecret)},
	})
} 

func GenerateToken(user_name string, id int) (string, error) {
	claims := jwt.MapClaims{
		"id": id,
		"user_name": user_name,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString([]byte(jwtSecret))
}


// func ParseToken(ctx *fiber.Ctx) {
// 	tokenString := ctx.Get("authorization")
// 	claims := jwt.MapClaims{}
// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(jwtSecret)
// 	})
// 	if err != nil {
// 		panic(err.Error())
// 	}


// }

