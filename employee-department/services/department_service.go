package services

import (
	"demo/employee-department/models"
	"demo/employee-department/repositories"
)

type DepartmentService interface{
	CreateDepartment(department *models.Department)error
	GetDepartments() ([]models.Department, error)
	GetByID(id uint)(*models.Department, error)
	UpdateDepartment(*models.Department)error
	DeleteDepartment(id uint)error
}

type departmentService struct{
	repo repositories.DepartmentRepository
}

func NewDepartmentService(repo repositories.DepartmentRepository) DepartmentService{
	return &departmentService{
		repo:repo,
	}
}

func(s *departmentService) CreateDepartment(department *models.Department) error{
	return s.repo.Create(department)
}

func(s *departmentService) GetDepartments()([]models.Department, error){
	return s.repo.GetAll()
}

func(s *departmentService) GetByID(id uint)(*models.Department, error){
	return s.repo.GetByID(id)
}

func(s *departmentService) UpdateDepartment(department *models.Department) error{
	return s.repo.Update(department)
}

func(s *departmentService) DeleteDepartment(id uint) error{
	return s.repo.Delete(id)
}