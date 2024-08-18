package handler 

import (
	"github.com/26thavenue/payroll_go/internal/services"
	"github.com/gofiber/fiber/v2"
)

type OvertimeHandler struct {
	OvertimeService *services.OvertimeService
}

func NewOvertimeHandler(OvertimeService *services.OvertimeService) *OvertimeHandler {
	return &OvertimeHandler{OvertimeService: OvertimeService}
}

func (h *OvertimeHandler) CreateOvertime (c *fiber.Ctx) error {

	return nil
}

func (h *OvertimeHandler) GetOvertimeByID (c *fiber.Ctx) error {

	return nil
}

func (h *OvertimeHandler) GetAllOvertimes (c *fiber.Ctx) error {

	return nil
}

func (h *OvertimeHandler) DeleteOvertime (c *fiber.Ctx) error {

	return nil
}

func (h *OvertimeHandler) UpdateOvertime (c *fiber.Ctx) error {

	return nil
}