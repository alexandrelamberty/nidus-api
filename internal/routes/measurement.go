package routes

import (
	"nidus-server/internal/handlers"
	"nidus-server/pkg/service"

	"github.com/gofiber/fiber/v2"
)

func MeasurementRouter(app fiber.Router, service service.MeasurementService) {
	app.Get("/Measurements", handlers.GetAllMeasurements(service))
	app.Post("/Measurements", handlers.CreateMeasurement(service))
	app.Get("/Measurements/:id", handlers.ReadMeasurement(service))
	app.Post("/Measurements/:id", handlers.UpdateMeasurement(service))
	app.Delete("/Measurements/:id", handlers.DeleteMeasurement(service))
}
