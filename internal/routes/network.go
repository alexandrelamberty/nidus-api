package routes

import (
	"nidus-server/internal/handlers"
	"nidus-server/pkg/service"

	"github.com/gofiber/fiber/v2"
)

func NetworkRouter(app fiber.Router, service service.NetworkService) {
	app.Get("/network/scan", handlers.Scan(service))
	app.Get("/network/ip/:mac", handlers.Ip(service))
}
