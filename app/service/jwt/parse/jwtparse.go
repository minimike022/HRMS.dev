package parsejwt

import (
	"github.com/gofiber/fiber/v2"
)

func ParseRefreshToken(ctx *fiber.Ctx) string{
	refresh_cookie := ctx.Cookies("refresh_token")

	return refresh_cookie
}

func ParseAccessToken(ctx *fiber.Ctx) string {
	access_cookie := ctx.Cookies("access_token")
	return access_cookie
}