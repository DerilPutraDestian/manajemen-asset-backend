package service

import (
	models "asset-management/model"
	"asset-management/repository"
)

type AssetCategoryService struct {
	repo *repository.CategoryRepository
}

func NewAssetCategoryService(repo *repository.CategoryRepository) *AssetCategoryService {
	return &AssetCategoryService{repo: repo}
}

func (s *AssetCategoryService) ListCategories(assetCode, search string, limit, offset int) ([]models.Category, int64, error) {
	return s.repo.GetAll(assetCode, search, limit, offset)
}

func (s *AssetCategoryService) GetCategory(id string) (*models.Category, error) {
	return s.repo.GetByID(id)
}

func (s *AssetCategoryService) CreateCategory(category *models.Category) error {
	return s.repo.Create(category)
}

func (s *AssetCategoryService) UpdateCategory(category *models.Category) error {
	return s.repo.Update(category)
}

func (s *AssetCategoryService) DeleteCategory(category *models.Category) error {
	return s.repo.Delete(category)
}
