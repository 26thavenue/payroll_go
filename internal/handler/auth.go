package handler

import (
	_ "encoding/json"
	"net/http"

	"github.com/26thavenue/payroll_go/internal/db"
	"github.com/26thavenue/payroll_go/internal/services"
	"github.com/26thavenue/payroll_go/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthService(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var creds struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

	if err := c.BodyParser(&creds); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	user, err := h.authService.Authenticate(creds.Username, creds.Password)

	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	accessToken, err := utils.GenerateAccessToken(user.ID, string(user.Role))
    if err != nil {
       return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate access token",
		})
    
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate refresh token"})
    }

	return c.JSON(fiber.Map{
        "access_token":  accessToken,
        "refresh_token": refreshToken,
    })
	
}

func (h *AuthHandler) Register (c *fiber.Ctx) error {

	var user db.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	err := h.authService.Register(&user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to register user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User registered successfully",
	})

}