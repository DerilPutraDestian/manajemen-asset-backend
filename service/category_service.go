package service

import (
	models "asset-management/model"
	"asset-management/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll(assetCode, search string, limit, offset int) ([]models.Category, int64, error) {
	return s.repo.GetAll(assetCode, search, limit, offset)
}

func (s *CategoryService) GetByID(id string) (*models.Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) Create(category *models.Category) error {
	return s.repo.Create(category)
}

func (s *CategoryService) Update(category *models.Category) error {
	return s.repo.Update(category)
}

func (s *CategoryService) Delete(category *models.Category) error {
	return s.repo.Delete(category)
}
