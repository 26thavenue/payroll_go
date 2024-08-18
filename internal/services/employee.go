package services

import (
	"github.com/26thavenue/payroll_go/internal/db"
	"github.com/26thavenue/payroll_go/internal/repository"
)

type EmployeeService struct {
	employeeRepo repository.EmployeeRepository
}

func NewEmployeeService(employeeRepo repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{employeeRepo: employeeRepo}
}

func (s *EmployeeService) CreateEmployee(employee *db.Employee) error{
	return s.employeeRepo.CreateEmployee(employee)
}

func (s *EmployeeService)GetEmployeeByID(id uint) (*db.Employee, error){
	return s.employeeRepo.GetEmployeeByID(id)
}


func (s *EmployeeService)ListEmployees(limit,offset int)([]*db.Employee, error){
	return s.employeeRepo.ListEmployees(limit,offset )
}

func (s *EmployeeService)UpdateEmployee(employee db.Employee) error{
	return s.employeeRepo.UpdateEmployee(&employee)
}

func (s *EmployeeService)DeleteEmployee(id uint) error{
	return s.employeeRepo.DeleteEmployee(id)
}