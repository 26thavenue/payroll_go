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
    FullName          string    `gorm:"size:100;not null"`   
    Employee Employee `gorm:"foreignKey:UserID"`
}

type Employee struct {
	ID            uint      `gorm:"primaryKey"`           
    Position      string    `gorm:"size:50;not null"`    
    Salary        float64   `gorm:"not null"`             
    Bonuses       float64   `gorm:"default:0"`            
    TaxID         string    `gorm:"size:50;unique"`      
    PhoneNumber   string    `gorm:"size:20"`            
    Address       string    `gorm:"size:255"`            
   	DepartmentID  uint        
    Department    *Department   `gorm:"foreignKey:DepartmentID"`            
    DateOfJoining *time.Time                             
    CreatedAt     time.Time                              
    UpdatedAt     time.Time
    UserID        uint          `gorm:"not null;constraint:OnDelete:CASCADE;"`
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
	ID            uint      `gorm:"primaryKey"`           
    EmployeeID    uint      `gorm:"not null"`            
    BasicSalary   float64   `gorm:"not null"`             
    Bonuses       float64   `gorm:"default:0"`            
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

