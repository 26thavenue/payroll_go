package handler

import (
	"github.com/26thavenue/payroll_go/internal/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"github.com/26thavenue/payroll_go/internal/db"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(UserService *services.UserService) *UserHandler {
	return &UserHandler{UserService: UserService}
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	user, err := h.UserService.GetUserByID(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(user)
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.UserService.ListUsers(1, 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(users)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	err = h.UserService.DeleteUser(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	var user db.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	user.ID = uint(userID)
	err = h.UserService.UpdateUser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *UserHandler) GetUserByEmail(c *fiber.Ctx) error {
	email := c.Query("email")
	if email == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Email is required")
	}

	user, err := h.UserService.FindByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(user)
}