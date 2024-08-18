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
    ListEmployees(limit,offset int) ([]*db.Employee, error)
}

// Implement the interface for a GORM-based repository
type GormEmployeeRepository struct {
    db *gorm.DB
}

func (r *GormEmployeeRepository) CreateEmployee(employee *db.Employee) error {
    return r.db.Create(employee).Error
}

func (r *GormEmployeeRepository) GetEmployeeByID(id uint) (*Employee, error) {
    var employee Employee
    err := r.db.First(&employee, id).Error //Why not find
    return &employee, err
}

func(r *GormEmployeeRepository)ListEmployees(limit,offset int)([]*db.Employee, error){
    var employees []*db.Employee

    err := r.db.Limit(limit).Offset(offset).Find(&employees).Error

    return employees, err
}


func (r *GormPayrollRepository)DeleteEmployee(id uint) error{
	return r.db.Delete(&db.Employee{}, id).Error
}

