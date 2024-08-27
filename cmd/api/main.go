package main

import (
	"fmt"

	"github.com/26thavenue/payroll_go/config"
	"github.com/26thavenue/payroll_go/db"
	"github.com/26thavenue/payroll_go/handlers"
	"github.com/26thavenue/payroll_go/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {

	envConfig := config.NewEnvConfig()

	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName: "Payroll Automation",
		ServerHeader: "Fiber",
	})

	// repositories
	payrollRepository := repositories.NewPayrollRepository(db)

	//Routing
	server := app.Group("/api")

	//handlers

	handlers.NewPayrollHandler(server.Group("/payroll"), payrollRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}