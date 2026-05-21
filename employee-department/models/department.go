package models

type Department struct {
	ID   uint   `json:"id" gorm:"primaryKey;"`
	Name string `json:"name" gorm:"type:varchar(20);not null"`

	Employees []Employee `json:"-" gorm:"foreignKey:DepartmentID"`
}
