package repository

import (
	"asset-management/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll(page, limit int) ([]models.Category, int64, error)
	FindByID(id int) (models.Category, error)
	Create(category *models.Category) error
	Update(category *models.Category) error
	Delete(id int) error
}

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepo{db}
}

func (r *categoryRepo) FindAll(page, limit int) ([]models.Category, int64, error) {
	var data []models.Category
	var total int64

	offset := (page - 1) * limit

	r.db.Model(&models.Category{}).Count(&total)
	err := r.db.Limit(limit).Offset(offset).Find(&data).Error

	return data, total, err
}

func (r *categoryRepo) FindByID(id int) (models.Category, error) {
	var data models.Category
	err := r.db.First(&data, id).Error
	return data, err
}

func (r *categoryRepo) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepo) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepo) Delete(id int) error {
	return r.db.Delete(&models.Category{}, id).Error
}
