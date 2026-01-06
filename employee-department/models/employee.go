package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model

	EmployeeName string `gorm:"type:varchar(100);not null" json:"employeeName"`
	Email        string `gorm:"type:varchar(100);not null;unique" json:"email"`

	DepartmentID uint   `gorm:"not null" json:"departmentId"`

}
