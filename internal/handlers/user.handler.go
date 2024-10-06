package handlers

import (
	"pos-gofiber/internal/domain/models"
	"pos-gofiber/internal/domain/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	user, err := h.userService.GetUser(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	return c.JSON(user)
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not retrieve users")
	}

	return c.JSON(users)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not hash password")
	}
	user.Password = string(hashedPassword)

	createdUser, err := h.userService.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not create user")
	}

	return c.Status(fiber.StatusCreated).JSON(createdUser)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}
	user.ID = uint(id)

	updatedUser, err := h.userService.UpdateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not update user")
	}

	return c.JSON(updatedUser)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	if err := h.userService.DeleteUser(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not delete user")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
