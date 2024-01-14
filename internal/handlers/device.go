package handlers

import (
	"net/http"
	"nidus-server/internal/requests"
	"nidus-server/internal/responses"
	"nidus-server/pkg/domain"
	"nidus-server/pkg/service"

	"github.com/gofiber/fiber/v2"
)

func PairDevice(service service.DeviceService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request requests.PairDeviceRequest
		err := c.BodyParser(&request)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}

		// Create the device we want to pair
		var device domain.Device
		device.Name = request.Name
		device.Ip = request.Ip
		device.Mac = request.Mac
		device.Paired = true
		result, err := service.CreateDevice(&device)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}

		return c.JSON(responses.DeviceSuccessResponse(result, "Device paired successfully"))
	}
}

func GetAllDevices(service service.DeviceService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.ListDevices()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}

		return c.JSON(responses.DevicesSuccessResponse(result, "Devices listed successfully"))
	}
}

// FIXME: need to be removed! Only pairing a device is allowed
func CreateDevice(service service.DeviceService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var device domain.Device
		err := c.BodyParser(&device)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}
		device.Paired = false

		result, err := service.CreateDevice(&device)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}

		return c.JSON(responses.DeviceSuccessResponse(result, "Device created successfully"))
	}
}

func ReadDevice(service service.DeviceService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		result, err := service.ReadDevice(id)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}

		return c.JSON(responses.DeviceSuccessResponse(result, "Device retrieved successfully"))
	}
}

func UpdateDevice(service service.DeviceService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		var requestBody domain.Device
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}

		result, err := service.UpdateDevice(id, &requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}

		return c.JSON(responses.DeviceSuccessResponse(result, "Device updated successfully"))
	}
}

func DeleteDevice(service service.DeviceService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		err := service.DeleteDevice(id)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}

		return c.JSON(responses.DeviceSuccessResponse(nil, "Device deleted successfully"))
	}
}
