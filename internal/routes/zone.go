package routes

import (
	"nidus-server/internal/handlers"
	"nidus-server/pkg/service"

	"github.com/gofiber/fiber/v2"
)

func ZoneRouter(app fiber.Router, service service.ZoneService) {
	app.Get("/zones", handlers.GetAllZones(service))
	app.Post("/zones", handlers.CreateZone(service))
	app.Get("/zones/:id", handlers.ReadZone(service))
	app.Post("/zones/:id", handlers.UpdateZone(service))
	app.Delete("/zones/:id", handlers.DeleteZone(service))
}
