package services

import (
	"demo/employee-department/models"
	"demo/employee-department/repositories"
)

type EmployeeService interface{
	CreateEmployee(employee *models.Employee) error
	GetEmployees() ([]models.Employee, error)
	GetEmployeeByID(id uint) (*models.Employee, error)
	UpdateEmployee(id uint, updates map[string]interface{}) error
	DeleteEmployee(id uint) error
}

type employeeService struct{
	repo repositories.EmployeeRepository
}

func NewEmployeeService(repo repositories.EmployeeRepository) EmployeeService{
	return &employeeService{
		repo: repo,
	}
}

func (s *employeeService) CreateEmployee(employee *models.Employee) error {
	return s.repo.Create(employee)
}

func (s *employeeService) GetEmployees() ([]models.Employee, error) {
	return s.repo.GetAll()
}

func (s *employeeService) GetEmployeeByID(id uint) (*models.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *employeeService) UpdateEmployee(id uint, updates map[string]interface{}) error {
	return s.repo.Update(id,updates)
}

func (s *employeeService) DeleteEmployee(id uint) error {
	return s.repo.Delete(id)
}
