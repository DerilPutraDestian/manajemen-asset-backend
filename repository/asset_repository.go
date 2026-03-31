package repository

import (
	"asset-management/models"

	"gorm.io/gorm"
)

type AssetRepository interface {
	FindAll(page, limit int) ([]models.Asset, int64, error)
	FindByID(id int) (models.Asset, error)
	Create(asset *models.Asset) error
	Update(asset *models.Asset) error
	Delete(id int) error
}

type repo struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) AssetRepository {
	return &repo{db}
}

func (r *repo) FindAll(page, limit int) ([]models.Asset, int64, error) {
	var assets []models.Asset
	var total int64

	offset := (page - 1) * limit
	r.db.Model(&models.Asset{}).Count(&total)

	err := r.db.Preload("Category").
		Limit(limit).
		Offset(offset).
		Find(&assets).Error

	return assets, total, err
}

func (r *repo) FindByID(id int) (models.Asset, error) {
	var asset models.Asset
	err := r.db.Preload("Category").First(&asset, id).Error
	return asset, err
}

func (r *repo) Create(asset *models.Asset) error {
	return r.db.Create(asset).Error
}

func (r *repo) Update(asset *models.Asset) error {
	return r.db.Save(asset).Error
}

func (r *repo) Delete(id int) error {
	return r.db.Delete(&models.Asset{}, id).Error
}
