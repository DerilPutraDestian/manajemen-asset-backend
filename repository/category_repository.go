package repository

import (
	models "asset-management/model"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll(assetCode, search string, limit, offset int) ([]models.Category, int64, error) {
	var data []models.Category
	var total int64

	query := r.db.Model(&models.Category{})

	if assetCode != "" {
		query = query.Where("asset_code = ?", assetCode)
	}

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	err := query.Count(&total).
		Limit(limit).
		Offset(offset).
		Find(&data).Error

	return data, total, err
}

func (r *CategoryRepository) GetByID(id string) (*models.Category, error) {
	var data models.Category
	err := r.db.First(&data, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *CategoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *CategoryRepository) Delete(category *models.Category) error {
	return r.db.Delete(category).Error
}
