package repositories

import (
	"demo/employee-department/models"
	"errors"

	"gorm.io/gorm"
)


type EmployeeRepository interface {
	Create(employee *models.Employee) error
	GetAll() ([]models.Employee, error)
	GetByID(id uint) (*models.Employee, error)
	Update(id uint,updates map[string]interface{}) error
	Delete(id uint) error
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{
		db: db,
	}
}

func (r *employeeRepository) Create(employee *models.Employee) error {
	return r.db.Create(employee).Error
}

func (r *employeeRepository) GetAll() ([]models.Employee, error) {

	var employees []models.Employee

	err := r.db.Preload("Department").Find(&employees).Error

	return employees, err
}

func (r *employeeRepository) GetByID(id uint) (*models.Employee, error) {

	var employee models.Employee

	err := r.db.Preload("Department").First(&employee, id).Error
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (r *employeeRepository) Update(id uint,updates map[string]interface{}) error {
	result := r.db.Model(&models.Employee{}).Where("id = ?",id).Updates(updates)

	if result.RowsAffected == 0{
		return errors.New("Employee not found")
	}
	return result.Error
}

func (r *employeeRepository) Delete(id uint) error {
	return r.db.Delete(&models.Employee{}, id).Error
}