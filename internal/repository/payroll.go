package repository

import (
	"github.com/26thavenue/payroll_go/internal/db"
	_"gorm.io/gorm"
)

type Payroll = db.Payroll

type PayrollRpository interface {
	CreatePayroll(payroll *Payroll) error
	GetPayrollById(id uint) (*Payroll, error)
	UpdatePayroll(payroll *Payroll) error
	ListPayroll()([]Payroll, error)
	DeletePayroll(id uint) error
}