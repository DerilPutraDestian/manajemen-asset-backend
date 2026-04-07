package repository

import (
	models "asset-management/model"

	"gorm.io/gorm"
)

type HistoryRepository interface {
	Create(h *models.AssetHistory) error
}

type historyRepo struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) HistoryRepository {
	return &historyRepo{db}
}

func (r *historyRepo) Create(h *models.AssetHistory) error {
	return r.db.Create(h).Error
}
