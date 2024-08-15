package services

import (
	"github.com/26thavenue/payroll_go/internal/db"
	"github.com/26thavenue/payroll_go/internal/repository"
)

type DepartmentService struct {
	deductionRepo repository.DepartmentRepository
}

func NewDepartmentService(deductionRepo repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{deductionRepo: deductionRepo}
}

func (s *DepartmentService) CreateDepartment(department db.Department) error{
	return s.deductionRepo.CreateDepartment(&department)
}

func (s *DepartmentService)GetByID(id uint) (*db.Department, error){
	return s.deductionRepo.GetDepartmentById(id)
}


func (s *DepartmentService)ListDepartments()([]db.Department, error){
	return s.deductionRepo.ListDepartment()
}


func (s *DepartmentService)DeleteDepartment(id uint) error{
	return s.deductionRepo.DeleteDepartment(id)
}