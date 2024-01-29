package handlers

import (
	"net/http"
	"nidus-server/internal/requests"
	"nidus-server/internal/responses"
	"nidus-server/pkg/domain"
	"nidus-server/pkg/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func PairDevice(service service.DeviceService, capabilityService service.CapabilityService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: Check pairing key is valid
		// pairingKey := c.Params("token")

		// Parse the request
		var request requests.PairDeviceRequest
		err := c.BodyParser(&request)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}

		// Validate the request
		validate := validator.New()
		if err := validate.Struct(request); err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}

		// Verify that the device has existing capabilities
		deviceCapabilityIDs, err := capabilityService.VerifyDeviceCapabilities(request.Capabilities)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}

		device := domain.Device{
			Name:          request.Name,
			Ip:            request.Ip,
			Mac:           request.Mac,
			Paired:        true,
			CapabilityIDs: *deviceCapabilityIDs,
		}

		result, err := service.CreateDevice(&device)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}

		// Loop through the capabilities and add the device
		// for _, capability := range capabilities {
		// 	err := capabilityService.AddDeviceToCapability(capability.ID, device.ID)
		// 	if err != nil {
		// 		c.Status(http.StatusInternalServerError)
		// 		return c.JSON(responses.DeviceErrorResponse(err.Error()))
		// 	}
		// }

		return c.JSON(responses.DeviceSuccessResponse(result, "Device paired successfully"))
	}
}

func SetupDevice(service service.DeviceService, capabilityService service.CapabilityService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		device_id := c.Params("id")

		// Parse the request
		var request requests.SetupDeviceRequest
		err := c.BodyParser(&request)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}

		// Validate the request
		validate := validator.New()
		if err := validate.Struct(request); err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}

		device := domain.Device{
			// CapabilityIDs: request.CapabilityIDs,
			ZoneID: request.ZoneID,
		}

		result, err := service.UpdateDevice(device_id, &device)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.DeviceErrorResponse(err.Error()))
		}

		return c.JSON(responses.DeviceSuccessResponse(result, "Device paired successfully"))
	}
}

// func UnPairDevice(service service.DeviceService) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		id := c.Params("id")

// 		err := service.UnPairDevice(id)
// 		if err != nil {
// 			c.Status(http.StatusInternalServerError)
// 			return c.JSON(responses.DeviceErrorResponse(err.Error()))
// 		}

// 		return c.JSON(responses.DeviceSuccessResponse(nil, "Device unpaired successfully"))
// 	}
// }

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
