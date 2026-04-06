package repository

import (
	models "asset-management/model"

	"gorm.io/gorm"
)

type AssetRepository struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) *AssetRepository {
	return &AssetRepository{db: db}
}

func (r *AssetRepository) GetAll(assetCode, search string, limit, offset int) ([]models.Asset, int64, error) {
	var assets []models.Asset
	var count int64

	query := r.db.Model(&models.Asset{})

	if assetCode != "" {
		query = query.Where("asset_code = ?", assetCode)
	}

	err := query.Count(&count).
		Preload("Category").
		Limit(limit).
		Offset(offset).
		Find(&assets).Error

	return assets, count, err
}

func (r *AssetRepository) GetByID(id string) (*models.Asset, error) {
	var asset models.Asset
	err := r.db.Preload("Category").
		Preload("Maintenances").
		First(&asset, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &asset, nil
}

func (r *AssetRepository) Create(asset *models.Asset) error {
	if err := r.db.Create(asset).Error; err != nil {
		return err
	}
	return r.db.Preload("Category").First(asset, "id = ?", asset.ID).Error
}
func (r *AssetRepository) Update(asset *models.Asset) error {
	return r.db.Save(asset).Error
}

func (r *AssetRepository) Delete(asset *models.Asset) error {
	return r.db.Delete(asset).Error
}
