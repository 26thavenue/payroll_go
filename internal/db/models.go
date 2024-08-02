package db

import (
	"time"
)

type SpecialCondition string

const (
    LEAVE     SpecialCondition = "LEAVE"
    SICK      SpecialCondition = "SICK"
    EMERGENCY SpecialCondition = "EMERGENCY"
	NONE 	  SpecialCondition = "NONE"
)

type Role string

const (
    RoleAdmin  Role = "ADMIN"
    RoleManager Role = "MANAGER"
    RoleEmployee Role = "EMPLOYEE"
)

type User struct {
    ID       uint   `gorm:"primaryKey"`   
    Email    string `gorm:"uniqueIndex;not null;unique"`
    Password string `gorm:"not null"`
    Role     Role   `gorm:"not null;default:EMPLOYEE"`
    FullName          string    `gorm:"size:100;not null"`    // Employee's full name
    Employee Employee `gorm:"foreignKey:UserID"`
}

type Employee struct {
	ID            uint      `gorm:"primaryKey"`           // Unique identifier for the employee
    Position      string    `gorm:"size:50;not null"`     // Job position or title
    Salary        float64   `gorm:"not null"`             // Base salary of the employee
    Bonuses       float64   `gorm:"default:0"`            // Any bonuses the employee is eligible for
    TaxID         string    `gorm:"size:50;unique"`       // Tax identification number
    PhoneNumber   string    `gorm:"size:20"`              // Contact phone number
    Address       string    `gorm:"size:255"`             // Physical address
   	DepartmentID  uint         // Foreign key to Department
    Department    *Department   `gorm:"foreignKey:DepartmentID"`             // Department in which the employee works
    DateOfJoining *time.Time                              // Date the employee joined the organization
    CreatedAt     time.Time                              // Timestamp for when the record was created
    UpdatedAt     time.Time
    UserID        uint
}

type Department struct {
	ID         uint      `gorm:"primaryKey"`
    Name       string    `gorm:"size:100;not null;unique"`
    HeadID     *uint     `gorm:"unique"` 
    Head       *Employee  `gorm:"foreignKey:HeadID"` 
    Employees  []Employee 
    CreatedAt  time.Time
    UpdatedAt  time.Time
}

type Payroll struct{
	ID            uint      `gorm:"primaryKey"`           // Unique identifier for the payroll record
    EmployeeID    uint      `gorm:"not null"`             // Foreign key referencing Employee
    BasicSalary   float64   `gorm:"not null"`             // Basic salary for the pay period
    Bonuses       float64   `gorm:"default:0"`            // Total bonuses for the pay period
    Deductions    float64   `gorm:"default:0"`            // Total deductions for the pay period
    Tax           float64   `gorm:"default:0"`            // Tax amount for the pay period
    NetSalary     float64   `gorm:"not null"`             // Net salary after deductions
    PaymentDate   time.Time `gorm:"not null"`             // Date when the payment was made
    Employee      Employee  `gorm:"foreignKey:EmployeeID"`// Association to Employee
    CreatedAt     time.Time                              // Timestamp for when the record was created
    UpdatedAt     time.Time  
}

type Overtime struct {
    ID          uint      `gorm:"primaryKey"`
    EmployeeID  uint      `gorm:"not null"`
    HoursWorked float64   `gorm:"not null"` // Number of overtime hours worked
    Date        time.Time `gorm:"not null"` // Date when the overtime was worked
    Employee    Employee  `gorm:"foreignKey:EmployeeID"` // Association to Employee
}

type Deductions struct {
    ID                 uint             `gorm:"primaryKey"`
    EmployeeID         uint             `gorm:"not null"`
    NumberOfDaysAbsent int              `gorm:"default:0"`  // Number of days absent
    DaysAbsent         int              `gorm:"default:0"`  // Number of days absent due to deductions
    SpecialCondition   SpecialCondition `gorm:"size:20"`    // Special conditions for deduction
    CreatedAt          time.Time
    UpdatedAt          time.Time
    Employee           Employee         `gorm:"foreignKey:EmployeeID"` // Association to Employee
}

