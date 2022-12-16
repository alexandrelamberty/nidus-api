package routes

import (
	"nidus-server/internal/handlers"
	"nidus-server/pkg/service"

	"github.com/gofiber/fiber/v2"
)

func SettingsRouter(app fiber.Router, service service.ZoneService) {
	app.Get("/settings", handlers.GetAllZones(service))
	app.Post("/settings", handlers.CreateZone(service))
	app.Get("/settings/:id", handlers.ReadZone(service))
}
