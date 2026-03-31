package repository

import (
	"asset-management/models"

	"gorm.io/gorm"
)

type HistoryRepository interface {
	Create(history *models.AssetHistory) error
}

type historyRepo struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) HistoryRepository {
	return &historyRepo{db}
}

func (r *historyRepo) Create(history *models.AssetHistory) error {
	return r.db.Create(history).Error
}
