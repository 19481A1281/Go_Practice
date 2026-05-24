package models

type Employee struct{
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(20);not null"`
	Email string `json:"email" gorm:"type:varchar(20);not null"`
	Mobile string `json:"mobile" gorm:"type:char(10);not null"`
	DepartmentID uint `json:"department" gorm:"column:department;not null"`

	Department Department `json:"-" gorm:"foreignKey:DepartmentID"`
}