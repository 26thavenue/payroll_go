package repository


import (
	"github.com/26thavenue/payroll_go/internal/db"
	_"gorm.io/gorm"
)

type Deductions = db.Deductions

type DeductionsRepository interface {
	CreateDeductions(deductions *Deductions ) error
	GetDeductionsById(id uint) (*Deductions , error)
	ListDeductions()([]Deductions , error)
	DeleteDeductions(id uint) error
}