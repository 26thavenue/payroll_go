package db

import (
    "fmt"
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	// "github.com/26thavenue/payroll_go/internal/db/models"
    "os"
)

var DB *gorm.DB


func InitDB() {
    var err error
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    
    DB.AutoMigrate(&Employee{}, &Payroll{}, &Department{}, &Deductions{}, &Overtime{})
}
