package validaterole

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)


func ValidateAdmin(ctx *fiber.Ctx) error{
	claims := ctx.Locals("claims").(jwt.MapClaims)
	role := claims["user_role"].(string)
	
	if role != "Admin" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized Access",
		})
	}

	return ctx.Next()
}

func ValidateHR(ctx *fiber.Ctx) error{
	claims := ctx.Locals("claims").(jwt.MapClaims)
	role := claims["user_role"].(string)
	fmt.Println(role)
	if role != "HR"{
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized Access",
		})
	}else if role != "Admin"{
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized Access",
		})
	}
	
	return ctx.Next()
}
