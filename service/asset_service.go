package service

import (
	"asset-management/models"
	"asset-management/repository"
	"asset-management/utils"
	"fmt"
)

type AssetService interface {
	GetAll(page, limit int) ([]models.Asset, int64, error)
	GetByID(id int) (models.Asset, error)
	Create(asset *models.Asset) error
	Update(id int, asset *models.Asset) error
	Delete(id int) error
}

type service struct {
	repo repository.AssetRepository
}

func NewAssetService(r repository.AssetRepository) AssetService {
	return &service{r}
}

func (s *service) GetAll(page, limit int) ([]models.Asset, int64, error) {
	return s.repo.FindAll(page, limit)
}

func (s *service) GetByID(id int) (models.Asset, error) {
	return s.repo.FindByID(id)
}

func (s *service) Create(asset *models.Asset) error {
	path := fmt.Sprintf("qrcodes/%s.png", asset.Code)
	if err := utils.GenerateQR(asset.Code, path); err != nil {
		return err
	}
	asset.QRCode = path
	return s.repo.Create(asset)
}

func (s *service) Update(id int, asset *models.Asset) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	asset.AssetID = existing.AssetID
	return s.repo.Update(asset)
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}
