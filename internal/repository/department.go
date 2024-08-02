package repository

import (
	"github.com/26thavenue/payroll_go/internal/db"
	_"gorm.io/gorm"
)

type Department = db.Department

type DepartmentRepository interface {
	CreateDepartment(department *Department) error
	GetDepartmentById(id uint) (*Department, error)
	ListDepartment()([]Department, error)
	DeleteDepartment(id uint) error
}