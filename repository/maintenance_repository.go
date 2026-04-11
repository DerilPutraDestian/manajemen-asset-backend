package repository

import (
	models "asset-management/model"

	"gorm.io/gorm"
)

type MaintenanceRepository struct {
	db *gorm.DB
}

func NewMaintenanceRepository(db *gorm.DB) *MaintenanceRepository {
	return &MaintenanceRepository{db: db}
}
func (r *MaintenanceRepository) GetAll(limit, offset int) ([]models.Maintenance, int64, error) {
	var data []models.Maintenance
	var total int64

	db := r.db.Model(&models.Maintenance{})

	// Hitung total data untuk keperluan pagination di frontend
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Ambil data dengan Preload Asset
	err := db.Preload("Asset").
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&data).Error

	return data, total, err
}
func (r *MaintenanceRepository) Create(m *models.Maintenance) error {
	// Omit Asset agar tidak mencoba insert ulang data master Asset
	return r.db.Omit("Asset").Create(m).Error
}

func (r *MaintenanceRepository) GetByID(id string) (*models.Maintenance, error) {
	var m models.Maintenance
	err := r.db.Preload("Asset").First(&m, "id = ?", id).Error
	return &m, err
}

func (r *MaintenanceRepository) GetByAssetID(assetID string) ([]models.Maintenance, error) {
	var data []models.Maintenance
	err := r.db.Where("asset_id = ?", assetID).Order("created_at DESC").Find(&data).Error
	return data, err
}

func (r *MaintenanceRepository) Update(m *models.Maintenance) error {
	return r.db.Model(m).
		Select("IssueDescription", "MaintenanceStatus", "StartDate", "EndDate", "UpdatedAt").
		Updates(m).Error
}
