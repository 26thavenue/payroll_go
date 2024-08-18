package repository

import (
	"github.com/26thavenue/payroll_go/internal/db"
	"gorm.io/gorm"
)

type Department = db.Department

type DepartmentRepository interface {
	CreateDepartment(department *Department) error
	GetDepartmentById(id uint) (*Department, error)
	ListDepartment(limit, offset int)([]*db.Department, error)
	DeleteDepartment(id uint) error
}

type GormDepartmentRepository struct {
	db *gorm.DB
}


func (r *GormDepartmentRepository) GetDepartmentById(id uint) (*Department, error) {
	var department Department
	err := r.db.First(&department, id).Error 
    return &department, err
}

func (r *GormDepartmentRepository) CreateDepartment(department *Department) error {
	return r.db.Create(department).Error
}

func (r *GormDepartmentRepository) DeleteDepartment(id *uint) error {
	var department Department

	return r.db.Delete(&department, id).Error
}

func (r *GormDepartmentRepository) ListDepartments(limit, offset int)([]*db.Department, error){
	var departments []*db.Department

	err := r.db.Limit(limit).Offset(offset).Find(&departments).Error

	return departments, err

}