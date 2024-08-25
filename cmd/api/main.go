package main

import (
	"github.com/26thavenue/payroll_go/handlers"
	"github.com/26thavenue/payroll_go/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Payroll Automation",
		ServerHeader: "Fiber",
	})

	// repositories
	payrollRepository := repositories.NewPayrollRepository(nil)

	//Routing
	server := app.Group("/api")

	//handlers

	handlers.NewPayrollHandler(server.Group("/payroll"), payrollRepository)

	app.Listen(":8080")
}