package location

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllLocations(ctx *fiber.Ctx) error {
	//TODO: connect to dynamodb and get all locations
	return ctx.Status(200).Send([]byte("Hello World"))
}

func GetLocationByCity(ctx *fiber.Ctx) {

}

func GetLocationByCityAndCounty(ctx *fiber.Ctx) {

}

func GetLocationByCityCountyAndCountry(ctx *fiber.Ctx) {

}

func GetLocationByCityPinCode(ctx *fiber.Ctx) {

}
