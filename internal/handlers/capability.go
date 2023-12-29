package handlers

import (
	"fmt"
	"net/http"
	"nidus-server/internal/responses"
	"nidus-server/pkg/domain"
	"nidus-server/pkg/service"

	"github.com/gofiber/fiber/v2"
)

// GetCapabilities is a function to get all services from the database
func GetAllCapabilities(service service.CapabilityService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("GetCapabilities")
		result, err := service.ListCapabilities()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.CapabilityErrorResponse(err.Error()))
		}
		return c.JSON(responses.CapabilitiesSuccessResponse(result, "OK"))
	}
}

func CreateCapability(service service.CapabilityService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody domain.Capability
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.CapabilityErrorResponse(err.Error()))
		}
		result, err := service.CreateCapability(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.CapabilityErrorResponse(err.Error()))
		}
		return c.JSON(responses.CapabilitySuccessResponse(result, "ok"))
	}
}

func ReadCapability(service service.CapabilityService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody domain.Capability
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.CapabilityErrorResponse(err.Error()))
		}
		result, err := service.ReadCapability("1")
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.CapabilityErrorResponse(err.Error()))
		}
		return c.JSON(responses.CapabilitySuccessResponse(result, "ok"))
	}
}

func UpdateCapability(service service.CapabilityService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody domain.Capability
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.CapabilityErrorResponse(err.Error()))
		}
		result, err := service.UpdateCapability(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.CapabilityErrorResponse(err.Error()))
		}
		return c.JSON(responses.CapabilitySuccessResponse(result, "ok"))
	}
}

func DeleteCapability(service service.CapabilityService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody domain.Capability
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.CapabilityErrorResponse(err.Error()))
		}
		result, err := service.UpdateCapability(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.CapabilityErrorResponse(err.Error()))
		}
		return c.JSON(responses.CapabilitySuccessResponse(result, "ok"))
	}
}
