package models

import (
	"context"
	"time"
)

type Payroll struct {
	ID           string
	EmployeeName string
	GrossSalary  uint
	Taxes        uint
	Date         time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PayrollRepository interface {
	GetAll(ctx context.Context) ([]*Payroll, error)
	GetOne(ctx context.Context, payrollID string) (*Payroll, error)
	CreatePayroll(ctx context.Context, payroll Payroll) (*Payroll, error)
	DeletePayroll(ctx context.Context,payrollID string ) error
}