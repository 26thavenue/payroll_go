package services

import (
	"github.com/26thavenue/payroll_go/internal/db"
	"github.com/26thavenue/payroll_go/internal/repository"
)

type OvertimeService struct {
	deductionRepo repository.OvertimeRepository
}

func NewOvertimeService(deductionRepo repository.OvertimeRepository) *OvertimeService {
	return &OvertimeService{deductionRepo: deductionRepo}
}

func (s *OvertimeService) CreateOvertime(Overtime db.Overtime) error{
	return s.deductionRepo.CreateOvertime(&Overtime)
}

func (s *OvertimeService)GetByID(id uint) (*db.Overtime, error){
	return s.deductionRepo.GetOvertimeById(id)
}


func (s *OvertimeService)ListOvertimes(limit, offset int)([]*db.Overtime, error){
	return s.deductionRepo.ListOvertime(limit, offset )
}


func (s *OvertimeService)DeleteOvertime(id uint) error{
	return s.deductionRepo.DeleteOvertime(id)
}