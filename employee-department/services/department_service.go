package services

import (
	"github.com/19481A1281/employee-department/models"
	"github.com/19481A1281/employee-department/repository"
)

type DepartmentService struct {
	departmentRepo repository.IDepartmentRepository
}

func NewDepartmentService(departmentRepo repository.IDepartmentRepository) *DepartmentService {
	return &DepartmentService{departmentRepo: departmentRepo}
}

func (s *DepartmentService) CreateDepartment(department *models.Department) error {
	return s.departmentRepo.CreateDepartment(department)
}

func (s *DepartmentService) GetAllDepartments() ([]models.Department, error) {
	return s.departmentRepo.GetAllDepartments()
}

func (s *DepartmentService) GetDepartmentByID(id uint) (*models.Department, error) {
	return s.departmentRepo.GetDepartmentByID(id)
}

func (s *DepartmentService) UpdateDepartment(department *models.Department) error {
	return s.departmentRepo.UpdateDepartment(department)
}

func (s *DepartmentService) DeleteDepartment(id uint) error {
	return s.departmentRepo.DeleteDepartment(id)
}
