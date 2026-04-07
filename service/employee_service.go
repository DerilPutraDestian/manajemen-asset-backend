package service

import (
	models "asset-management/model"
	"asset-management/repository"
)

type EmployeeService struct {
	repo *repository.EmployeeRepository
}

func NewEmployeeService(repo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) ListEmployees(search string, limit, offset int) ([]models.Employee, int64, error) {
	return s.repo.GetAll(search, limit, offset)
}

func (s *EmployeeService) GetEmployee(id uint) (*models.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *EmployeeService) CreateEmployee(e *models.Employee) error {
	return s.repo.Create(e)
}

func (s *EmployeeService) UpdateEmployee(e *models.Employee) error {
	return s.repo.Update(e)
}

func (s *EmployeeService) DeleteEmployee(id uint) error {
	return s.repo.Delete(id)
}
