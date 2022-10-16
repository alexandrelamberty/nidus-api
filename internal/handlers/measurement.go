package handlers

import (
	"fmt"
	"net/http"
	"nidus-server/internal/requests"
	"nidus-server/internal/responses"
	"nidus-server/pkg/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllMeasurements(service service.MeasurementService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.ListMeasurements()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.MeasurementErrorResponse(err.Error()))
		}
		return c.JSON(responses.ListMeasurementSuccessResponse(result, "OK"))
	}
}

func CreateMeasurement(service service.MeasurementService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("Measurement::Post", c.IP())
		var requestBody requests.CreateMeasurementRequest
		requestBody.Timestamp = primitive.NewDateTimeFromTime(time.Now())
		err := c.BodyParser(&requestBody)
		fmt.Println(requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.MeasurementErrorResponse(err.Error()))
		}
		// Fixme request to domain ?
		result, err := service.CreateMeasurement(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.MeasurementErrorResponse(err.Error()))
		}
		return c.JSON(responses.CreateMeasurementSuccessResponse(result, "ok"))
	}
}

func ReadMeasurement(service service.MeasurementService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var deviceId = c.Params("id")
		var timestamp = c.Query("date")
		var sensorType = c.Query("type")

		if deviceId == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.MeasurementErrorResponse("No ID"))
		}
		result, err := service.ReadMeasurement(deviceId, sensorType, timestamp)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.MeasurementErrorResponse(err.Error()))
		}
		return c.JSON(responses.ReadMeasurementSuccessResponse(result, "ok"))
	}
}
