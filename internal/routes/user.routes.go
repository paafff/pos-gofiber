package routes

import (
	"pos-gofiber/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router, userHandler *handlers.UserHandler) {
	router.Get("/user/:id", userHandler.GetUser)
	router.Post("/user", userHandler.CreateUser)
	router.Put("/user/:id", userHandler.UpdateUser)
	router.Delete("/user/:id", userHandler.DeleteUser)
	router.Get("/users", userHandler.GetUsers)
}
