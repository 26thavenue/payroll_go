package handler

import (
	"strconv"

	"github.com/26thavenue/payroll_go/internal/db"
	"github.com/26thavenue/payroll_go/internal/services"
	"github.com/gofiber/fiber/v2"
)

type EmployeeHandler struct {
	EmployeeService *services.EmployeeService
}

func NewEmployeeHandler(employeeService *services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{EmployeeService: employeeService}
}

func (h *EmployeeHandler) CreateEmployee (c *fiber.Ctx) error {
	var employee db.Employee

	if err := c.BodyParser(&employee); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
    }

	if err := h.EmployeeService.CreateEmployee(&employee); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusCreated).JSON(employee)

	
}

func (h *EmployeeHandler) GetEmployeeByID (c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil{
		return c.Status(fiber.StatusBadRequest).SendString("error: Invalid Request")
	}

	user, err := h.EmployeeService.GetEmployeeByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(user)
}

func (h *EmployeeHandler) GetAllEmployees (c *fiber.Ctx) error {
	
	employees, err := h.EmployeeService.ListEmployees(1,10)

	if err !=nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(employees)
}

func (h *EmployeeHandler) DeleteEmployee (c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil{
		return c.Status(fiber.StatusBadRequest).SendString("error: Invalid Request")
	}

	err = h.EmployeeService.DeleteEmployee(uint(id))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusNoContent).SendString("Employee Deleted Succesfully")
	
}

func (h *EmployeeHandler) UpdateEmployee (c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid employee ID"})
    }

    var employee db.Employee
    if err := c.BodyParser(&employee); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
    }
	employee.ID = uint(id)
    if err := h.EmployeeService.UpdateEmployee(employee); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(employee)
	
}