package repository

import (
	"github.com/26thavenue/payroll_go/internal/db"
	"gorm.io/gorm"
)

type Deductions = db.Deductions

type DeductionsRepository interface {
	CreateDeductions(deductions *Deductions ) error
	GetDeductionsById(id uint) (*Deductions , error)
	ListDeductions(limit,offset int)([]*db.Deductions , error)
	DeleteDeductions(id uint) error
}

type GormDeductionsRepository struct {
	db *gorm.DB
}


func (r *GormDeductionsRepository) GetDeductionsById(id uint) (*Deductions , error) {
	var deductions Deductions
	err := r.db.First(&deductions, id).Error 
    return &deductions, err
}

func (r *GormDeductionsRepository) CreateDeductions(deductions *Deductions ) error {
	return r.db.Create(deductions).Error
}

func (r *GormDeductionsRepository)DeleteDeduction(id uint)error{
	return r.db.Delete(&db.Deductions{}, id).Error
}

func (r *GormDeductionsRepository) ListDeductions(limit, offset int)([]*db.Deductions, error){
	var deductions []*db.Deductions

	err := r.db.Limit(limit).Offset(offset).Find(&deductions).Error
	return deductions,err
}