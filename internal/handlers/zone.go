package handlers

import (
	"net/http"
	"nidus-server/internal/responses"
	"nidus-server/pkg/domain"
	"nidus-server/pkg/service"

	"github.com/gofiber/fiber/v2"
)

// GetZones is a function to get all services from the database
func GetAllZones(service service.ZoneService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.ListZones()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ZoneErrorResponse(err.Error()))
		}
		return c.JSON(responses.ZonesSuccessResponse(result, "OK"))
	}
}

func CreateZone(service service.ZoneService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody domain.Zone
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ZoneErrorResponse(err.Error()))
		}
		result, err := service.CreateZone(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ZoneErrorResponse(err.Error()))
		}
		return c.JSON(responses.ZoneSuccessResponse(result, "ok"))
	}
}

func ReadZone(service service.ZoneService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody domain.Zone
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ZoneErrorResponse(err.Error()))
		}
		result, err := service.ReadZone("1")
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ZoneErrorResponse(err.Error()))
		}
		return c.JSON(responses.ZoneSuccessResponse(result, "ok"))
	}
}

func UpdateZone(service service.ZoneService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody domain.Zone
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ZoneErrorResponse(err.Error()))
		}
		result, err := service.UpdateZone(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ZoneErrorResponse(err.Error()))
		}
		return c.JSON(responses.ZoneSuccessResponse(result, "ok"))
	}
}

func DeleteZone(service service.ZoneService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody domain.Zone
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.ZoneErrorResponse(err.Error()))
		}
		result, err := service.UpdateZone(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.ZoneErrorResponse(err.Error()))
		}
		return c.JSON(responses.ZoneSuccessResponse(result, "ok"))
	}
}
