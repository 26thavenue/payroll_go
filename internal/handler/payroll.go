package handler 

import (
	"github.com/26thavenue/payroll_go/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/26thavenue/payroll_go/internal/db"
)

type PayrollHandler struct {
	PayrollService *services.PayrollService
}

func NewPayrollHandler(PayrollService *services.PayrollService) *PayrollHandler {
	return &PayrollHandler{PayrollService: PayrollService}
}

func (h *PayrollHandler) CreatePayroll (c *fiber.Ctx) error {

	var payroll db.Payroll

	if err := c.BodyParser(&payroll);err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	err := h.PayrollService.CreatePayroll(payroll)

	if err!= nil{

	}



	return nil
}

func (h *PayrollHandler) GetPayrollByID (c *fiber.Ctx) error {

	return nil
}

func (h *PayrollHandler) GetAllPayrolls (c *fiber.Ctx) error {

	return nil
}

func (h *PayrollHandler) DeletePayroll (c *fiber.Ctx) error {

	return nil
}

func (h *PayrollHandler) UpdatePayroll (c *fiber.Ctx) error {

	return nil
}