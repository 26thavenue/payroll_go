package services

import (
	"github.com/26thavenue/payroll_go/internal/db"
	"github.com/26thavenue/payroll_go/internal/repository"
)

type PayrollService struct {
	payrollRepo repository.PayrollRepository
}

func NewPayrollService(payrollRepo repository.PayrollRepository) *PayrollService {
	return &PayrollService{payrollRepo: payrollRepo}
}

func (s *PayrollService) CreatePayroll(payroll db.Payroll) error{
	return s.payrollRepo.CreatePayroll(&payroll)
}

func (s *PayrollService)GetPayrollById(id uint) (*db.Payroll, error){
	return s.payrollRepo.GetPayrollById(id)
}


func (s *PayrollService)ListPayroll()([]db.Payroll, error){
	return s.payrollRepo.ListPayroll()
}

func (s *PayrollService)UpdatePayroll(payroll db.Payroll) error{
	return s.payrollRepo.UpdatePayroll(&payroll)
}

func (s *PayrollService)DeletePayroll(id uint) error{
	return s.payrollRepo.DeletePayroll(id)
}