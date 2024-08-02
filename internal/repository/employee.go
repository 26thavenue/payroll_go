package repository

import (
	"github.com/26thavenue/payroll_go/internal/db"
	"gorm.io/gorm"
)

type Employee = db.Employee

//Define an interface for Employee repository
type EmployeeRepository interface {
    CreateEmployee(employee *Employee) error
    GetEmployeeByID(id uint) (*Employee, error)
    UpdateEmployee(employee *Employee) error
    DeleteEmployee(id uint) error
    ListEmployees() ([]Employee, error)
}

// Implement the interface for a GORM-based repository
type GormEmployeeRepository struct {
    db *gorm.DB
}

func (r *GormEmployeeRepository) CreateEmployee(employee *Employee) error {
    return r.db.Create(employee).Error
}

func (r *GormEmployeeRepository) GetEmployeeByID(id uint) (*Employee, error) {
    var employee Employee
    err := r.db.First(&employee, id).Error
    return &employee, err
}


