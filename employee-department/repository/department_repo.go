package repository

import (

	"github.com/19481A1281/employee-department/models"
	"gorm.io/gorm"
)

type IDepartmentRepository interface {
	CreateDepartment(department *models.Department) error
	GetDepartmentByID(id uint) (*models.Department, error)
	GetAllDepartments() ([]models.Department, error)
	UpdateDepartment(department *models.Department) error
	DeleteDepartment(id uint) error
}


type DepartmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) IDepartmentRepository {
	return &DepartmentRepository{db: db}
}

func (r *DepartmentRepository) CreateDepartment(department *models.Department) error {
	return r.db.Create(department).Error
}

func (r *DepartmentRepository) GetAllDepartments() ([]models.Department, error) {
	var departments []models.Department
	err := r.db.Preload("Employees").Find(&departments).Error
	return departments, err
}

func (r *DepartmentRepository) GetDepartmentByID(id uint) (*models.Department, error) {
	var department models.Department
	err := r.db.Preload("Employees").First(&department, id).Error
	return &department, err
}

func (r *DepartmentRepository) UpdateDepartment(department *models.Department) error {
	return r.db.Save(department).Error
}

func (r *DepartmentRepository) DeleteDepartment(id uint) error {
	return r.db.Delete(&models.Department{}, id).Error
}





