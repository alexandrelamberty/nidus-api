package routes

import (
	"nidus-server/internal/handlers"
	"nidus-server/pkg/service"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service service.UserService) {
	app.Get("/users", handlers.GetAllUsers(service))
	app.Post("/users", handlers.CreateUser(service))
	app.Get("/users/:id", handlers.ReadUser(service))
	app.Post("/users/:id", handlers.UpdateUser(service))
	app.Delete("/users/:id", handlers.DeleteUser(service))
}
