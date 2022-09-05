package routes

import (
	"nidus-server/internal/handlers"
	"nidus-server/pkg/service"

	"github.com/gofiber/fiber/v2"
)

func CapabilityRouter(app fiber.Router, service service.CapabilityService) {
	app.Get("/capabilities", handlers.GetAllCapabilities(service))
	app.Post("/capabilities", handlers.CreateCapability(service))
	app.Get("/capabilities/:id", handlers.ReadCapability(service))
	app.Post("/capabilities/:id", handlers.UpdateCapability(service))
	app.Delete("/capabilities/:id", handlers.DeleteCapability(service))
}
