package repository

import (
	"github.com/26thavenue/payroll_go/internal/db"
	"gorm.io/gorm"
)

type Payroll = db.Payroll

type PayrollRepository interface {
	CreatePayroll(payroll *Payroll) error
	GetPayrollById(id uint) (*Payroll, error)
	UpdatePayroll(payroll *Payroll) error
	ListPayroll(limit, offset int)([]*db.Payroll, error)
	DeletePayroll(id uint) error
}

type GormPayrollRepository struct {
    db *gorm.DB
}

func (r *GormPayrollRepository) CreatePayroll(payroll *Payroll) error {
    return r.db.Create(payroll).Error
}

func (r *GormPayrollRepository) GetPayrollById(id uint) (*Payroll, error) {
    var payroll Payroll
    err := r.db.First(&payroll, id).Error 
    return &payroll, err
}

func (r *GormPayrollRepository) ListPayroll(limit, offset int) ([]*db.Payroll, error) {
	var payrolls []*db.Payroll
	err := r.db.Limit(limit).Offset(offset).Find(&payrolls).Error
	return payrolls, err
}

func (r *GormPayrollRepository)DeletePayroll(id uint) error{
	return r.db.Delete(&db.Payroll{}, id).Error
}
