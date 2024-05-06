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
	refresh_auth := parsejwt.ParseRefreshToken(ctx)

	refresh_token, err := jwt.ParseWithClaims(refresh_auth, jwt.MapClaims{}, func(t *jwt.Token)(interface{}, error) {
		return []byte(os.Getenv("REFRESH_TOKEN_KEY")), nil
	})

	claims := refresh_token.Claims.(jwt.MapClaims)
	refresh_exp_time := claims["exp"].(float64)
	refresh_exp_timef := time.Unix(int64(refresh_exp_time),0)

	if time.Until(refresh_exp_timef) <= 0 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Error" : "Token Expired",
		})
	}

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	ctx.Locals("claims", claims)

	return ctx.Next()
}

func ValidateAccessToken(ctx *fiber.Ctx) error {
	//Fetch Token from Parsing Cookies
	access_auth := parsejwt.ParseAccessToken(ctx)

	//Parse token with payloads
	access_token, err := jwt.ParseWithClaims(access_auth, jwt.MapClaims{}, func(t *jwt.Token)(interface{},error) {
		return []byte(os.Getenv("ACCESS_TOKEN_KEY")), nil
	})

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	claims := access_token.Claims.(jwt.MapClaims)

	//Fetch access_exp from jwt claims
	access_exp_time := claims["exp"].(float64)

	//Convert float64 to UNIX TIME FORMAT
	access_exp_timef := time.Unix(int64(access_exp_time),0)

	//Validate if access token is expired
	if time.Until(access_exp_timef) <= 0 {
		RenewAccessToken(ctx)
	}
	
	return ctx.Next()
}

func RenewAccessToken(ctx *fiber.Ctx) error {
	//Get the refresh token cookie
	refresh_auth := parsejwt.ParseRefreshToken(ctx)
	//Parse refresh token
	refresh_token, _ := jwt.ParseWithClaims(refresh_auth, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("REFRESH_TOKEN_KEY"), nil
	})

	claims := refresh_token.Claims.(jwt.MapClaims)

	//Fetched claims data from refresh_token
	id := claims["id"].(float64)
	user_name:= claims["user_name"].(string)
	user_role := claims["user_role"].(string)

	//Fetch expiration time from jwtclaims
	refresh_exp := claims["exp"].(float64)
	//Convert Float64 to UNIX TIME FORMAT
	refresh_exp_unix := time.Unix(int64(refresh_exp), 0)

	//Validate if token is not expired
	if time.Until(refresh_exp_unix) <= 0 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error" : "Token Expired",
		})
	}

	//Generate  Access Token
	access_token, _ := jwtgenerate.GenerateAccessToken(user_name, int(id), user_role)

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