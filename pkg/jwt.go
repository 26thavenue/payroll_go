package pkg


import (
    "github.com/26thavenue/payroll_go/internal/db"
    "github.com/gofiber/fiber/v2"
    jwtware "github.com/gofiber/jwt/v3"
    "github.com/golang-jwt/jwt/v4"
	"github.com/26thavenue/payroll_go/internal/config"
)

var SECRET_KEY = config.JWT_SECRET

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
    user, ok := c.Locals("user").(*jwt.Token)

    if !ok {
        return nil, ErrNoUserInContext
    }
    claims, ok := user.Claims.(jwt.MapClaims)

    if !ok {
        return nil, ErrInvalidTokenClaims
    }

     idClaim, ok := claims["id"]
    if !ok {
        return nil, ErrMissingIDClaim
    }
    idFloat, ok := idClaim.(float64)
    if !ok {
        return nil, ErrInvalidIDType
    }

    roleClaim, ok := claims["role"]
    if !ok {
        return nil, ErrMissingRoleClaim
    }
    
    roleString, ok := roleClaim.(string)
    if !ok {
        return nil, ErrInvalidRoleType
    }

    
    return &db.User{
        ID:   uint(idFloat),
        Role: db.Role(roleString),
    }, nil
}