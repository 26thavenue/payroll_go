package repository

import (
	"github.com/26thavenue/payroll_go/internal/db"
	"gorm.io/gorm"
)

type Overtime = db.Overtime

type OvertimeRepository interface {
	CreateOvertime(overtime *Overtime) error
	GetOvertimeById(id uint) (*Overtime, error)
	UpdateOvertime(overtime *Overtime) error
	ListOvertime(limit, offset int)([]*db.Overtime, error)
	DeleteOvertime(id uint) error
}

type GormOvertimeRepository struct {
    db *gorm.DB
}

func (r *GormOvertimeRepository) CreateOvertime(overtime *Overtime) error {
    return r.db.Create(overtime).Error
}

func (r *GormOvertimeRepository) GetOvertimeById(id uint) (*Overtime, error) {
    var overtime Overtime
    err := r.db.First(&overtime, id).Error 
    return &overtime, err
}

func (r *GormOvertimeRepository) ListOvertime(limit, offset int) ([]*db.Overtime, error) {
	var overtimes []*db.Overtime
	err := r.db.Limit(limit).Offset(offset).Find(&overtimes).Error
	return overtimes, err
}

func (r *GormOvertimeRepository)DeleteOvertime(id uint) error{
	return r.db.Delete(&db.Overtime{}, id).Error
}