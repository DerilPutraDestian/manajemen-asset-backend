package service

import (
	"asset-management/models"
	"asset-management/repository"
)

type CategoryService interface {
	GetAll(page, limit int) ([]models.Category, int64, error)
	GetByID(id int) (models.Category, error)
	Create(category *models.Category) error
	Update(id int, category *models.Category) error
	Delete(id int) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(r repository.CategoryRepository) CategoryService {
	return &categoryService{r}
}

func (s *categoryService) GetAll(page, limit int) ([]models.Category, int64, error) {
	return s.repo.FindAll(page, limit)
}

func (s *categoryService) GetByID(id int) (models.Category, error) {
	return s.repo.FindByID(id)
}

func (s *categoryService) Create(category *models.Category) error {
	return s.repo.Create(category)
}

func (s *categoryService) Update(id int, category *models.Category) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	category.CategoryID = existing.CategoryID
	return s.repo.Update(category)
}

func (s *categoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
