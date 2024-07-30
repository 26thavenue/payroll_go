package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main(){
	app := fiber.New()

	app.Use(logger.New(logger.Config{
    	Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))


	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Theree!")
	})

	app.Listen(":3000")
}