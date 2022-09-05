package handlers

import (
	"net/http"
	"nidus-server/internal/responses"
	"nidus-server/pkg/domain"
	"nidus-server/pkg/service"

	"github.com/gofiber/fiber/v2"
)

// GetUsers is a function to get all services from the database
func GetAllUsers(service service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.ListUsers()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}
		return c.JSON(responses.UsersSuccessResponse(result, "OK"))
	}
}

func CreateUser(service service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody domain.User
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}
		result, err := service.CreateUser(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}
		return c.JSON(responses.UserSuccessResponse(result, "ok"))
	}
}

func ReadUser(service service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		result, err := service.ReadUser(id)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}
		return c.JSON(responses.UserSuccessResponse(result, "ok"))
	}
}

func UpdateUser(service service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody domain.User
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}
		result, err := service.UpdateUser(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}
		return c.JSON(responses.UserSuccessResponse(result, "ok"))
	}
}

func DeleteUser(service service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := service.DeleteUser(id)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(responses.UserErrorResponse(err.Error()))
		}
		return c.JSON(responses.UserSuccessResponse(nil, "ok"))
	}
}
