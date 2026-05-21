package repositories

import (
	"demo/employee-department/models"

	"gorm.io/gorm"
)

type DepartmentRepository interface{
	Create(dept *models.Department) error
	GetAll() ([]models.Department, error)
	GetByID(id uint) (*models.Department, error)
}

type departmentRepository struct{
	db *gorm.DB
}

func NewDepartmentRepositoty(db *gorm.DB) DepartmentRepository{
	return &departmentRepository{db: db}
}

func (r *departmentRepository) Create(dept *models.Department) error{
	return r.db.Create(dept).Error
}

func (r *departmentRepository) GetAll() ([]models.Department, error) {
	var depts []models.Department
	err := r.db.Find(&depts).Error
	return depts, err
}

func (r *departmentRepository) GetByID(id uint) (*models.Department, error) {
	var dept models.Department
	err := r.db.First(&dept, id).Error
	if err != nil {
		return nil, err
	}
	return &dept, nil
}