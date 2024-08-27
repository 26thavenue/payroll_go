package db

import (
	"github.com/26thavenue/payroll_go/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.Payroll{})
}