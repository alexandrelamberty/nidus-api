package handlers

import (
	"nidus-server/pkg/service"

	"github.com/gofiber/fiber/v2"
)

func Scan(service service.NetworkService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result := service.Scan()
		// if err != nil {
		// 	c.Status(http.StatusInternalServerError)
		// 	return c.JSON(responses.ZoneErrorResponse(err.Error()))
		// }
		// return c.JSON(responses.ZonesSuccessResponse(result, "OK"))
		return c.JSON(result)
	}
}

func Ip(service service.NetworkService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrive path parameter
		mac := c.Params("mac")
		result := service.GetIp(mac)
		// if err != nil {
		// 	c.Status(http.StatusInternalServerError)
		// 	return c.JSON(responses.ZoneErrorResponse(err.Error()))
		// }
		// return c.JSON(responses.ZonesSuccessResponse(result, "OK"))
		return c.JSON(result)
	}
}
