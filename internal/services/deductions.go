package services

import (
	"github.com/26thavenue/payroll_go/internal/db"
	_ "github.com/26thavenue/payroll_go/internal/db"
	"github.com/26thavenue/payroll_go/internal/repository"
)

type DeductionService struct {
	deductionRepo repository.DeductionsRepository
}

func NewDeductionService(deductionRepo repository.DeductionsRepository) *DeductionService {
	return &DeductionService{deductionRepo: deductionRepo}
}

func (s *DeductionService) CreateDeductions(Deductions db.Deductions) error{
	return s.deductionRepo.CreateDeductions(&Deductions)
}

func (s *DeductionService)GetDeductionsByID(id uint) (*db.Deductions, error){
	return s.deductionRepo.GetDeductionsById(id)
}


func (s *DeductionService)ListDeductionss()([]db.Deductions, error){
	return s.deductionRepo.ListDeductions()
}


func (s *DeductionService)DeleteDeductions(id uint) error{
	return s.deductionRepo.DeleteDeductions(id)
}