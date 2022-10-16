package routes

import (
	"nidus-server/internal/handlers"
	"nidus-server/pkg/service"

	"github.com/gofiber/fiber/v2"
)

func DeviceRouter(app fiber.Router, service service.DeviceService) {
	app.Post("/devices/pair", handlers.PairDevice(service))
	app.Get("/devices", handlers.GetAllDevices(service))
	app.Post("/devices", handlers.CreateDevice(service))
	app.Get("/devices/:id", handlers.ReadDevice(service))
	app.Put("/devices/:id", handlers.UpdateDevice(service))
	app.Delete("/devices/:id", handlers.DeleteDevice(service))
}
