package db

import (
    "fmt"
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	// "github.com/26thavenue/payroll_go/internal/db/models"
    "github.com/26thavenue/payroll_go/internal/config"
)

var DB *gorm.DB

var DB_HOST = config.DB_HOST
var DB_USER = config.DB_USER
var	DB_PASSWORD = config.DB_PASSWORD
var	DB_NAME = config.DB_NAME
var	DB_PORT = config.DB_PORT


func InitDB() {
    var err error
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
        DB_HOST,
		DB_USER,
		DB_PASSWORD,
		DB_NAME,
		DB_PORT,
    )
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    
    DB.AutoMigrate(&Employee{}, &Payroll{}, &Department{}, &Deductions{}, &Overtime{}, &User{})
}
