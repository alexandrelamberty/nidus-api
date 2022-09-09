package handlers

import (
	"net/http"
	"nidus-server/internal/responses"
	"nidus-server/pkg/domain"
	"nidus-server/pkg/service"

	"github.com/gofiber/fiber/v2"
)

// GetMeasurements is a function to get all services from the database
func GetAllMeasurements(service service.MeasurementService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.ListMeasurements()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.MeasurementErrorResponse(err.Error()))
		}
		return c.JSON(responses.MeasurementsSuccessResponse(result, "OK"))
	}
}

func CreateMeasurement(service service.MeasurementService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody domain.Measurement
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.MeasurementErrorResponse(err.Error()))
		}
		result, err := service.CreateMeasurement(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.MeasurementErrorResponse(err.Error()))
		}
		return c.JSON(responses.MeasurementsuccessResponse(result, "ok"))
	}
}

func ReadMeasurement(service service.MeasurementService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody domain.Measurement
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.MeasurementErrorResponse(err.Error()))
		}
		result, err := service.ReadMeasurement("1")
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.MeasurementErrorResponse(err.Error()))
		}
		return c.JSON(responses.MeasurementsuccessResponse(result, "ok"))
	}
}

func UpdateMeasurement(service service.MeasurementService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody domain.Measurement
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.MeasurementErrorResponse(err.Error()))
		}
		result, err := service.UpdateMeasurement(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.MeasurementErrorResponse(err.Error()))
		}
		return c.JSON(responses.MeasurementsuccessResponse(result, "ok"))
	}
}

func DeleteMeasurement(service service.MeasurementService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody domain.Measurement
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.MeasurementErrorResponse(err.Error()))
		}
		result, err := service.UpdateMeasurement(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.MeasurementErrorResponse(err.Error()))
		}
		return c.JSON(responses.MeasurementsuccessResponse(result, "ok"))
	}
}
