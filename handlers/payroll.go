package handlers

import (
	"context"
	"time"

	"github.com/26thavenue/payroll_go/models"
	"github.com/gofiber/fiber/v2"
)


type PayrollHandler struct {
	repository models.PayrollRepository
}

func (h *PayrollHandler)GetAll(ctx *fiber.Ctx) error{
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5 * time.Second))
	defer cancel()

	payroll, err := h.repository.GetAll(context)

	if err != nil{
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":"failed",
			"message": err.Error(),
		})
	}

	return  ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status":"success",
			"message": err.Error(),
			"data":payroll,
		})
}

func (h *PayrollHandler)GetOne(ctx *fiber.Ctx, payrollId string)error{
	return nil
}
func (h *PayrollHandler)CreatePayroll(ctx *fiber.Ctx )error{
	return nil
}
func (h *PayrollHandler)DeletePayroll(ctx *fiber.Ctx, payrollId string)error{
	return nil
}


func NewPayrollHandler (router fiber.Router,repository models.PayrollRepository){
	handler := &PayrollHandler{
		repository: repository,
	}

	router.Get("/", handler.GetAll)
	// router.Get("/:payrollId", handler.GetOne)
	router.Post("/", handler.CreatePayroll)
	// router.Delete("/", handler.DeletePayroll)
}