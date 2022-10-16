package routes

import (
	"nidus-server/internal/handlers"
	"nidus-server/pkg/service"

	"github.com/gofiber/fiber/v2"
)

func MeasurementRouter(app fiber.Router, service service.MeasurementService) {
	app.Get("/measurements", handlers.GetAllMeasurements(service))
	app.Post("/measurements", handlers.CreateMeasurement(service))
	app.Get("/measurements/:id", handlers.ReadMeasurement(service))
}
