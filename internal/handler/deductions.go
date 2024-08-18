package handler

import (
	"github.com/26thavenue/payroll_go/internal/services"
	"github.com/gofiber/fiber/v2"
)

type DeductionsHandler struct {
	DeductionsService *services.DeductionService
}

func NewDeductionsHandler(DeductionsService *services.DeductionService) *DeductionsHandler {
	return &DeductionsHandler{DeductionsService: DeductionsService}
}

func (h *DeductionsHandler) CreateDeductions (c *fiber.Ctx) error {

	return nil
}

func (h *DeductionsHandler) GetDeductionsByID (c *fiber.Ctx) error {

	return nil
}

func (h *DeductionsHandler) GetAllDeductionss (c *fiber.Ctx) error {

	return nil
}

func (h *DeductionsHandler) DeleteDeductions (c *fiber.Ctx) error {

	return nil
}

func (h *DeductionsHandler) UpdateDeductions (c *fiber.Ctx) error {

	return nil
}