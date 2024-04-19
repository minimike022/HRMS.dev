package routesData

import (
	"github.com/gofiber/fiber/v2"
)

type TryData struct {
	Name string
	Age int
}

func GetData(ctx *fiber.Ctx) error {
	msg := "Hello world"
	return ctx.Status(fiber.StatusOK).JSON(msg)
}

func PostData(ctx *fiber.Ctx) error {
	try := new(TryData)
	err := ctx.BodyParser(try)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		return err
	}

	data := TryData{
		Name: try.Name,
		Age: try.Age,
	}

	return ctx.Status(fiber.StatusOK).JSON(data)

}


func PutData(ctx *fiber.Ctx) error {
	msg := "Hello World"

	return ctx.Status(fiber.StatusOK).JSON(msg)
}
