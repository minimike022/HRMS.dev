package jwt

import (
	"os"
	"time"
	//jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// func JwtAuth() func(ctx *fiber.Ctx) error {
// 	return jwtware.New(jwtware.Config{
// 		ErrorHandler: func (ctx *fiber.Ctx, err error) error {
// 			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": "Unauthorized Access",
// 			})
// 		},
// 		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("SECRET_KEY"))},
// 	})
// } 

func ValidateToken(ctx *fiber.Ctx) {
	auth := ParseToken(ctx)

	token, err := jwt.ParseWithClaims(auth, jwt.MapClaims{}, func(t *jwt.Token)(interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		panic(err.Error())
	}

	

}

func GenerateToken(user_name string, id int) (string, error) {
	claims := jwt.MapClaims{
		"id": id,
		"user_name": user_name,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}


func ParseToken(ctx *fiber.Ctx) string {
	tokenHeader := ctx.Get(fiber.HeaderAuthorization)

	tokenString := tokenHeader[len("Bearer"):]

	return tokenString
}

