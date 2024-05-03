package jwtvalidate

import (
	jwtgenerate "hrms-api/app/service/jwt"
	parsejwt "hrms-api/app/service/jwt/parse"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateRefreshToken(ctx *fiber.Ctx) error {
	auth := parsejwt.ParseRefreshToken(ctx)

	_, err := jwt.ParseWithClaims(auth, jwt.MapClaims{}, func(t *jwt.Token)(interface{}, error) {
		return []byte(os.Getenv("REFRESH_TOKEN_KEY")), nil
	})

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	return ctx.Next()
}

func ValidateAccessToken(ctx *fiber.Ctx) error {
	auth := parsejwt.ParseAccessToken(ctx)

	token, err := jwt.ParseWithClaims(auth, jwt.MapClaims{}, func(t *jwt.Token)(interface{},error) {
		return []byte(os.Getenv("ACCESS_TOKEN_KEY")), nil
	})

	if err != nil {
		ctx.Status(fiber.StatusUnauthorized)
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	// expiration_time := claims["exp"].(float64)

	// if time.Unix(int64(expiration_time), 0) > time.Now()

	RenewAccessToken(ctx)
	
	ctx.Locals("claims", claims)

	return ctx.Next()
}

func RenewAccessToken(ctx *fiber.Ctx) error {
	//Get the refresh token cookie
	auth := parsejwt.ParseRefreshToken(ctx)


	//Parse refresh token
	token, _ := jwt.ParseWithClaims(auth, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("REFRESH_TOKEN_KEY"), nil
	})

	claims := token.Claims.(jwt.MapClaims)

	id := claims["id"].(float64)
	user_name:= claims["user_name"].(string)
	user_role := claims["user_role"].(string)

	access_token, err := jwtgenerate.GenerateAccessToken(user_name, int(id), user_role)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map {
			"error" : "Invalid Token",
		})
	}

	cookie := fiber.Cookie {
		Name: "access_token",
		Value: access_token,
		Expires: time.Now().Add(time.Minute * 15),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map {
		"message" : "Access Token Renewed",
	})
}