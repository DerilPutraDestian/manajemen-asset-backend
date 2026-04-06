package service

import (
	models "asset-management/model"
	"asset-management/repository"
)

type AssetService struct {
	repo *repository.AssetRepository
}

func NewAssetService(repo *repository.AssetRepository) *AssetService {
	return &AssetService{repo: repo}
}

func (s *AssetService) ListAssets(assetCode, search string, limit, offset int) ([]models.Asset, int64, error) {
	return s.repo.GetAll(assetCode, search, limit, offset)
}

func (s *AssetService) GetAsset(id string) (*models.Asset, error) {
	return s.repo.GetByID(id)
}

func (s *AssetService) CreateAsset(asset *models.Asset) error {
	return s.repo.Create(asset)
}

func (s *AssetService) UpdateAsset(asset *models.Asset) error {
	return s.repo.Update(asset)
}

func (s *AssetService) DeleteAsset(asset *models.Asset) error {
	return s.repo.Delete(asset)
}
