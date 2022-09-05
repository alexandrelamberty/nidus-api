package handlers

import (
	"net/http"
	"nidus-server/internal/responses"
	"nidus-server/pkg/domain"
	"nidus-server/pkg/service"

	"github.com/gofiber/fiber/v2"
)

// GetDevices is a function to get all services from the database
func GetAllDevices(service service.DeviceService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.ListDevices()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}
		return c.JSON(responses.DevicesSuccessResponse(result, "OK"))
	}
}

func CreateDevice(service service.DeviceService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var device domain.Device
		err := c.BodyParser(&device)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}
		result, err := service.CreateDevice(&device)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}
		return c.JSON(responses.DeviceSuccessResponse(result, "ok"))
	}
}

func ReadDevice(service service.DeviceService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.DeviceErrorResponse("no id"))
		}
		result, err := service.ReadDevice(id)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}
		return c.JSON(responses.DeviceSuccessResponse(result, "Device found"))
	}
}

func UpdateDevice(service service.DeviceService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody domain.Device
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}
		result, err := service.UpdateDevice(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}
		return c.JSON(responses.DeviceSuccessResponse(result, "ok"))
	}
}

func DeleteDevice(service service.DeviceService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.DeviceErrorResponse("Id not provided"))
		}
		err := service.DeleteDevice(id)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}
		return c.JSON(responses.DeviceSuccessResponse(nil, "Delete success"))
	}
}
