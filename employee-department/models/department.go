package models

import (
	"gorm.io/gorm"
)

type Department struct {
	gorm.Model

	DepartmentName string `gorm:"type:varchar(100);not null" json:"departmentName"`

	Employees []Employee `gorm:"foreignKey:DepartmentID" json:"employees"`

}
