package pkg


import (
    "github.com/26thavenue/payroll_go/internal/db"
    "github.com/gofiber/fiber/v2"
    jwtware "github.com/gofiber/jwt/v3"
    "github.com/golang-jwt/jwt/v4"
	"github.com/26thavenue/payroll_go/internal/config"
)

const SECRET_KEY = config.JWT_SECRET

func JWTAuthentication() fiber.Handler {
    return jwtware.New(jwtware.Config{
        SigningKey: []byte(SECRET_KEY),
        ErrorHandler: func(c *fiber.Ctx, err error) error {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Unauthorized",
            })
        },
    })
}

func ExtractUserFromToken(c *fiber.Ctx) (*db.User, error) {
    user := c.Locals("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    
    return &db.User{
        ID:   uint(claims["id"].(float64)),
        Role: db.Role(claims["role"].(string)),
    }, nil
}