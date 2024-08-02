package repository

import (
	"github.com/26thavenue/payroll_go/internal/db"
	_"gorm.io/gorm"
)

type Overtime = db.Overtime

type OvertimeRepository interface {
	CreateOvertime(overtime *Overtime) error
	GetOvertimeById(id uint) (*Overtime, error)
	UpdateOvertime(overtime *Overtime) error
	ListOvertime()([]Overtime, error)
	DeleteOvertime(id uint) error
}