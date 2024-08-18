package handler 

import (
	"github.com/26thavenue/payroll_go/internal/services"
	"github.com/gofiber/fiber/v2"
)

type DepartmentHandler struct {
	DepartmentService *services.DepartmentService
}

func NewDepartmentHandler(DepartmentService *services.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{DepartmentService: DepartmentService}
}

func (h *DepartmentHandler) CreateDepartment (c *fiber.Ctx) error {

	return nil
}

func (h *DepartmentHandler) GetDepartmentByID (c *fiber.Ctx) error {

	return nil
}

func (h *DepartmentHandler) GetAllDepartments (c *fiber.Ctx) error {

	return nil
}

func (h *DepartmentHandler) DeleteDepartment (c *fiber.Ctx) error {

	return nil
}

func (h *DepartmentHandler) UpdateDepartment (c *fiber.Ctx) error {

	return nil
}